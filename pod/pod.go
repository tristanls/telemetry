package pod

import (
	"github.com/tristanls/telemetry"
	"github.com/tristanls/telemetry/logger"
	"github.com/tristanls/telemetry/metrics"
)

// Pod Content is a bundle of Telemetry-related functionality.
type Content struct {
	Telemetry *telemetry.Telemetry
	Emitter   *telemetry.Emitter
	Log       *logger.Logger
	Metrics   *metrics.Metrics
}

// Constructs a new Telemetry Pod.
func New(telemetry *telemetry.Telemetry, emitter *telemetry.Emitter, log *logger.Logger, metrics *metrics.Metrics) *Content {
	return &Content{
		Telemetry: telemetry,
		Emitter:   emitter,
		Log:       log,
		Metrics:   metrics,
	}
}
