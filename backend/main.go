package main

import (
	"context"
	"errors"
	"flag"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"git.mstar.dev/mstar/goutils/other"
	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"

	"github.com/WNTR-Software/foxcal/backend/config"
	"github.com/WNTR-Software/foxcal/backend/storage"
	"github.com/WNTR-Software/foxcal/backend/storage/dbgen"
	"github.com/WNTR-Software/foxcal/backend/storage/models"
	"github.com/WNTR-Software/foxcal/backend/web"
)

/*
Some packages to consider:
caldav: https://pkg.go.dev/github.com/emersion/go-webdav/caldav
oidc: https://github.com/coreos/go-oidc
*/

func main() {
	other.SetupFlags() // To include the loglevel cli flag
	flag.Parse()

	// Configure writing logs to a persistent file
	// and configure logging based on the flag -loglevel
	// Example:
	// -loglevel debug;gorm=warn;web=debug
	// Global log level will be set to debug
	// Log level for gorm will be set to warn
	// Log level for web will be set to debug
	logfile := getLogFilePathOrNil()
	var logLevels map[string]zerolog.Level
	var logfileWriter io.Writer
	// Configure logging before writing any messages. Otherwise it would default to json for stderr
	if logfile == nil {
		logLevels = other.ConfigureLogging(logfileWriter)
		log.Warn().Msg("Can't write to target logfile, not creating one")
	} else {
		logfileWriter = &lumberjack.Logger{
			Filename:   *logfile,
			MaxSize:    500, // Megabytes
			MaxBackups: 3,
			MaxAge:     3, // Days
			Compress:   false,
		}
		logLevels = other.ConfigureLogging(logfileWriter)
		log.Info().Str("logfile", *logfile).Msg("Logging to stderr and logfile")
	}

	if err := other.LoadConfigOrWriteDefault(&config.Global, &config.DefaultConfig, *flagConfigFile); err != nil {
		if errors.Is(err, other.ErrFileDidntExist) {
			config.Global = config.DefaultConfig
			log.Info().
				Str("filename", *flagConfigFile).
				Msg("config file didn't exist, created it with default values")
		} else {
			log.Fatal().Str("filename", *flagConfigFile).Msg("Failed to load or write config file")
		}
	}

	// Setup the db, only sqlite supported for now
	// TODO: Add postgres once decision on read replica support is in
	log.Info().Msg("Configuring database")
	if !config.Global.Db.UseSqlite {
		log.Fatal().
			Msg("Postgres support has not yet been added, please use sqlite instead for now")
	}

	// Configure logging for db
	// Looks big since some work has to be done to skip past
	// the code in gorm and gorm-gen
	gormLogger := other.NewLogger(logfileWriter).Level(logLevels["gorm"])
	// Trace call location for all db calls
	gormLogger = gormLogger.Hook(
		zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
			if !(level == zerolog.DebugLevel || level == zerolog.TraceLevel) {
				return
			}
			skipCount := 8
			for i := 8; i < 15; i++ {
				_, file, _, ok := runtime.Caller(i)
				if !ok {
					break
				}
				if strings.Contains(file, "gorm") || strings.Contains(file, "dbgen") {
					continue
				}
				skipCount = i
				break
			}
			e.Caller(skipCount)
		}),
	)
	// Actual db setup is here
	gormdb, err := gorm.Open(sqlite.Open(config.Global.Db.SqliteFile), &gorm.Config{
		Logger: storage.NewGormLogger(gormLogger),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to db")
	}
	if err = gormdb.AutoMigrate(models.AllModels...); err != nil {
		log.Fatal().Err(err).Msg("Failed to automigrate database")
	}
	dbgen.SetDefault(gormdb)

	// Catch interrupt signal (ctrl-c) for clean shutdown
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	// And start the server
	server := web.NewServer(log.Level(logLevels["web"]))
	log.Info().Str("address", config.Global.BindAddress).Msg("Starting webserver")
	go func() {
		// Server runs in a separate goroutine to allow for a clean shutdown
		if err := server.Run(config.Global.BindAddress); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server failed to run")
		}
	}()
	// Since the main routine needs to wait for an interrupt signal
	// and then stop the server
	<-interruptChan
	log.Info().Msg("Shutting down")
	if err = server.Stop(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to cleanly stop the server")
	}
	log.Info().Msg("Shutdown successful")
	os.Exit(0)
}
