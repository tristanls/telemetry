# Telemetry

_Stability: 1 - [Experimental](https://github.com/tristanls/stability-index#stability-1---experimental)_

Telemetry is a helper for creating and emitting structured telemetry events like logs, metrics, or usage.

## Contributors

[@tristanls](https://github.com/tristanls)

## Contents

* [Usage](#usage)
    * [Logger](#logger)
    * [Metrics](#metrics)
* [Documentation](#documentation)
* [Releases](#releases)

## Usage

To run the below example run: `$ go run examples/json_to_stdout.go`.

```go
package main

import (
	"github.com/tristanls/telemetry"
)

var emitter = telemetry.NewEmitter()
var writer = telemetry.NewWriter()

func main() {

	emitter.AddListener(func(event *telemetry.Event) {
		writer.Write(event.Marshal())
	})

	event := telemetry.New().WithProvenance(telemetry.Fields{
		"import":  "github.com/tristanls/telemetry",
		"version": "0.0.0",
	})

	emitter.Emit(event.WithFields(telemetry.Fields{
		"type":    "log",
		"level":   "info",
		"message": "hello o/",
	}))

	detailedProvenance := event.WithProvenance(telemetry.Fields{
		"file": "json_to_stdout.go",
	})

	emitter.Emit(detailedProvenance.WithFields(telemetry.Fields{
		"type":        "metric",
		"name":        "web requests",
		"target_type": "counter",
		"unit":        "Req",
		"value":       1,
	}))

	emitter.Emit(detailedProvenance.WithFields(telemetry.Fields{
		"type":     "usage",
		"tenantId": "tristan1234",
		"usage": telemetry.Fields{
			"storage": telemetry.Fields{
				"request": telemetry.Fields{
					"unit":  "Req",
					"value": 2,
				},
			},
		},
	}))
}
```

Which should result in something like:

```json
{"level":"info","message":"hello o/","provenance":[{"import":"github.com/tristanls/telemetry","version":"0.0.0"}],"timestamp":"2017-02-18T22:02:35.452Z","type":"log"}
{"name":"web requests","provenance":[{"import":"github.com/tristanls/telemetry","version":"0.0.0"},{"file":"json_to_stdout.go"}],"target_type":"counter","timestamp":"2017-02-18T22:02:35.452Z","type":"metric","unit":"Req","value":1}
{"provenance":[{"import":"github.com/tristanls/telemetry","version":"0.0.0"},{"file":"json_to_stdout.go"}],"tenantId":"tristan1234","timestamp":"2017-02-18T22:02:35.452Z","type":"usage","usage":{"storage":{"request":{"unit":"Req","value":2}}}}
```

### Logger

Logger is a helper for emitting log events using Telemetry.

To run the example below run: `$ go run examples/logger.go`.

```go
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
```

### Metrics

Metrics is a helper for emitting metric events using Telemetry.

To run the example below run: `$ go run examples/metrics.go`.

```go
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
```

## Documentation

Please refer to [generated Go documentation](https://godoc.org/github.com/tristanls/telemetry)

## Releases

We follow semantic versioning policy ([semver.org](http://semver.org/)) with a caveat:

> Given a version number MAJOR.MINOR.PATCH, increment the:
>
>MAJOR version when you make incompatible API changes,<br/>
>MINOR version when you add functionality in a backwards-compatible manner, and<br/>
>PATCH version when you make backwards-compatible bug fixes.

**caveat**: Major version zero is a special case indicating development version that may make incompatible API changes without incrementing MAJOR version.