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

	// Prints out:
	// $ go run examples/json_to_stdout.go
	// {"level":"info","message":"hello o/","provenance":[{"import":"github.com/tristanls/telemetry","version":"0.0.0"}],"timestamp":"2017-02-18T22:02:35.452Z","type":"log"}
	// {"name":"web requests","provenance":[{"import":"github.com/tristanls/telemetry","version":"0.0.0"},{"file":"json_to_stdout.go"}],"target_type":"counter","timestamp":"2017-02-18T22:02:35.452Z","type":"metric","unit":"Req","value":1}
	// {"provenance":[{"import":"github.com/tristanls/telemetry","version":"0.0.0"},{"file":"json_to_stdout.go"}],"tenantId":"tristan1234","timestamp":"2017-02-18T22:02:35.452Z","type":"usage","usage":{"storage":{"request":{"unit":"Req","value":2}}}}
}
