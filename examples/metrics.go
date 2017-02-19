package main

import (
	"github.com/tristanls/telemetry"
	"github.com/tristanls/telemetry/metrics"
)

func main() {

	_telemetry := telemetry.New()
	emitter := telemetry.NewEmitter()
	writer := telemetry.NewWriter()

	emitter.AddListener(func(event *telemetry.Event) {
		writer.Write(event.WithProvenance(telemetry.Fields{
			"example": "github.com/tristanls/telemetry/examples/metrics.go",
		}).Marshal())
	})

	m := metrics.NewMetrics(_telemetry, emitter)

	m.Metric(_telemetry.WithFields(telemetry.Fields{
		"name":        "my.gauge",
		"target_type": "gauge",
		"unit":        "ms",
		"value":       133,
	}))
	m.Counter("my.counter", _telemetry.WithFields(telemetry.Fields{
		"unit":  "Req",
		"value": 123,
	}))
	m.Gauge("my.gauge", _telemetry.WithFields(telemetry.Fields{
		"unit":  "ms",
		"value": 133,
	}))
	m.Histogram("search.results.returned", _telemetry.WithFields(telemetry.Fields{
		"value": telemetry.Fields{
			"measureUnit":       "request",
			"sampleSizeUnit":    "Req",
			"updateCount":       2,
			"max":               7122,
			"mean":              4174.5,
			"median":            7122,
			"min":               1227,
			"percentile75":      7122,
			"percentile95":      7122,
			"percentile98":      7122,
			"percentile99":      7122,
			"percentile999":     7122,
			"standardDeviation": 2947.5,
			"sampleSize":        2,
		},
	}))
	m.Meter("requests", _telemetry.WithFields(telemetry.Fields{
		"value": telemetry.Fields{
			"rateUnit":          "Req/s",
			"updateCount":       11,
			"updateCountUnit":   "Req",
			"meanRate":          692.3612402089173,
			"oneMinuteRate":     0,
			"fiveMinuteRate":    0,
			"fifteenMinuteRate": 0,
		},
	}))
	m.Timer("request.latency", _telemetry.WithFields(telemetry.Fields{
		"value": telemetry.Fields{
			"measureUnit":       "ms",
			"rateUnit":          "Req/s",
			"sampleSizeUnit":    "Req",
			"updateCount":       2,
			"meanRate":          125.82722724393419,
			"oneMinuteRate":     0,
			"fiveMinuteRate":    0,
			"fifteenMinuteRate": 0,
			"max":               178,
			"mean":              89.01773188379212,
			"median":            178,
			"min":               0.03412902355194092,
			"percentile75":      178,
			"percentile95":      178,
			"percentile98":      178,
			"percentile99":      178,
			"percentile999":     178,
			"standardDeviation": 88.98293548572138,
			"sampleSize":        2,
		},
	}))
}
