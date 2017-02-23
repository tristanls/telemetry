package telemetry

import (
	"github.com/tristanls/telemetry/logger"
)

// Pod is a bundle of Telemetry-related functionality.
type Pod struct {
	Telemetry *Telemetry
	Emitter   *Emitter
	Log       *logger.Logger
}

// Constructs a new Telemetry Pod.
func NewPod(telemetry *Telemetry, emitter *Emitter, log *logger.Logger) *Pod {
	return &Pod{
		Telemetry: telemetry,
		Emitter:   emitter,
		Log:       log,
	}
}
