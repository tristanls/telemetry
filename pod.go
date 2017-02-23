package telemetry

import (
	"github.com/tristanls/telemetry/logger"
	"github.com/tristanls/telemetry/metrics"
)

// Pod is a bundle of Telemetry-related functionality.
type Pod struct {
	Telemetry *Telemetry
	Emitter   *Emitter
	Log       *logger.Logger
	Metrics   *metrics.Metrics
}

// Constructs a new Telemetry Pod.
func NewPod(telemetry *Telemetry, emitter *Emitter, log *logger.Logger, metrics *metrics.Metrics) *Pod {
	return &Pod{
		Telemetry: telemetry,
		Emitter:   emitter,
		Log:       log,
		Metrics:   metrics,
	}
}
