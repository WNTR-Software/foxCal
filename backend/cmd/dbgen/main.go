package main

import (
	"errors"
	"flag"

	"git.mstar.dev/mstar/goutils/other"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"

	"WNTR-Software/foxCal/backend/config"
	"WNTR-Software/foxCal/backend/storage/models"
)

func main() {
	other.SetupFlags()
	configFileName := flag.String("config", "config.toml", "Set the config file to use")
	outPath := flag.String("output", "./generated", "Set the dir to write output though")
	flag.Parse()
	other.ConfigureLogging(nil)
	err := other.LoadConfigOrWriteDefault(&config.Global, &config.DefaultConfig, *configFileName)
	if err != nil {
		if errors.Is(err, other.ErrFileDidntExist) {
			config.Global = config.DefaultConfig
			log.Info().
				Str("filename", *configFileName).
				Msg("config file didn't exist, created it with default values")
		} else {
			log.Fatal().Str("filename", *configFileName).Msg("Failed to load or write config file")
		}
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: *outPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// Init the db, required to know which syntax to use
	gormdb, _ := gorm.Open(postgres.Open(config.Global.Db.BuildPostgresDSN()), &gorm.Config{})
	g.UseDB(gormdb)

	// Apply the basic operations
	g.ApplyBasic(models.AllModels...)

	// Then link the custom queries to their relevant models
	g.ApplyInterface(func(models.ILink) {}, models.Link{})
	g.ApplyInterface(func(models.IUpcomingEvent) {}, models.UpcomingEvent{})

	// And build
	g.Execute()
}
