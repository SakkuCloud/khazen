package main

import (
	"flag"
	"github.com/evalphobia/logrus_sentry"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"khazen/app"
	"khazen/config"
	"os"
	"time"
)

func main() {
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	var configFile = flag.String("c", "/etc/khazen/config.yml", "config file location (must be yaml format)")
	flag.Parse()

	// log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if *debugMode {
		log.Infoln("Log in debug mode!")
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// config file
	if *debugMode {
		err := configor.New(&configor.Config{ENVPrefix: "KHAZEN", Verbose: true}).Load(&config.Config, *configFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := configor.New(&configor.Config{ENVPrefix: "KHAZEN"}).Load(&config.Config, *configFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// log file
	logFile, err := os.OpenFile(config.Config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Infof("Logging to %s", config.Config.LogFile)
		log.SetOutput(logFile)
	} else {
		log.Info("Failed to log to file, using default")
	}

	// sentry
	if config.Config.SentryDSN != "" {
		hook, _ := logrus_sentry.NewSentryHook(config.Config.SentryDSN, []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		})
		hook.Timeout = config.SentryTimeout * time.Second
		log.AddHook(hook)
	}

	// run app
	api := &app.App{}
	api.Init()
	api.Run(":" + config.Config.Port)
}
