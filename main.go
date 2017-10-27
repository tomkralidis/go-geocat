package main

import (
    "os"

    "github.com/sirupsen/logrus"
    "github.com/tomkralidis/go-geocat/config"
)

const Version = "0.1.0"

func main() {
    // get configuration
    cfg := config.GetConfig("config.yml")

    // setup logging
    log := initLog(cfg)

    log.Info("Starting go-geocat server Version " + Version)
    log.Info("Server URL: " + cfg.Server.Url)
    return
}


func initLog(cfg config.Config) logrus.Logger {
    var log = logrus.Logger{
        Out: os.Stderr,
        Formatter: new(logrus.TextFormatter),
        Hooks: make(logrus.LevelHooks),
        Level: logrus.DebugLevel,
    }
    return log
}