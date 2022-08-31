package main

import (
    "github.com/sirupsen/logrus"
    prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
     log.Formatter = new(prefixed.TextFormatter)
     log.Level = logrus.DebugLevel
}


func main() {

    // Samples 
    log.Info("Temperature changes")
    log.Debug("Temperature changes")
    log.Warning("Temperature changes")
    log.Error("Error Temperature changes")
    log.Trace("Ntarse Temperature changes")

    // Заканчивается программа
    // log.Panic("Panic Temperature changes")

    // log.WithFields(logrus.Fields{"prefix": "main",   "animal": "walrus","number": 8}).Debug("Started observing beach")
    // log.WithFields(logrus.Fields{"prefix": "sensor", "temperature": -4}).Info("Temperature changes")
}
