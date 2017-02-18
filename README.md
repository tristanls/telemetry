# Telemetry

_Stability: 1 - [Experimental](https://github.com/tristanls/stability-index#stability-1---experimental)_

Telemetry is a helper for creating and emitting structured telemetry events like logs, metrics, or usage.

## Contributors

[@tristanls](https://github.com/tristanls)

## Contents

* [Usage](#usage)
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