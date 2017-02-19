package main

import (
	"github.com/tristanls/telemetry"
	"github.com/tristanls/telemetry/logger"
)

func main() {

	_telemetry := telemetry.New()
	emitter := telemetry.NewEmitter()
	writer := telemetry.NewWriter()

	emitter.AddListener(func(event *telemetry.Event) {
		writer.Write(event.WithProvenance(telemetry.Fields{
			"example": "github.com/tristanls/telemetry/logger/logger.go",
		}).Marshal())
	})

	log := logger.NewLogger(_telemetry, emitter)

	log.Log(logger.Info, "hello")
	log.Logf(logger.Debug, "hello %v", "o/")
	log.Debug("debugging")
	log.Debugf("formatted %s", "debugging")
	log.Info("informational")
	log.Infof("informational with %s", "format")
	log.Warn("warning")
	log.Warnf("warning with %s", "format")
	log.Error("error happened")
	log.Errorf("error happened with %s", "format")
	log.Fatal("fatally oh, noes")
	log.Fatalf("fatal, but it is not a logger's job to %s", "stop the process :/")

	event := _telemetry.WithFields(telemetry.Fields{
		"these":    "will be",
		"included": "in all events extend this one",
	})

	log.Loge(logger.Warn, event, "hello")
	log.Logef(logger.Error, event, "formatted with %v", "provided format")
	log.Debuge(event, "debugged")
	log.Debugef(event, "formatted %s", "debugging")
	log.Infoe(event, "informational")
	log.Infoef(event, "informational with %s", "format")
	log.Warne(event, "warning")
	log.Warnef(event, "warning with %s", "format")
	log.Errore(event, "error happened")
	log.Erroref(event, "error happened with %s", "format")
	log.Fatale(event, "fatally oh, noes")
	log.Fatalef(event, "fatal, but it is not a logger's job to %s", "stop the process :/")
}
