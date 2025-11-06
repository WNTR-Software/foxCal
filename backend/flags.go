package main

import (
	"flag"
	"path"

	"golang.org/x/sys/unix"
)

var (
	flagConfigFile = flag.String("config", "config.toml", "Set the config file to use")
	flagLogFile    = flag.String(
		"logfile",
		"/var/log/foxcal/log",
		"Where to write logs to in addition to stderr",
	)
)

func getLogFilePathOrNil() *string {
	logfile := *flagLogFile
	logPath := path.Dir(logfile)
	if unix.Access(logPath, unix.W_OK) == nil {
		return &logfile
	} else {
		return nil
	}
}
