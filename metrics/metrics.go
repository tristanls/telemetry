package metrics

import (
	"sync"

	"github.com/tristanls/telemetry"
)

// Creates a new Metrics instance that will emit Events on provided telemetry emitter using provided telemetry
// configuration.
func NewMetrics(telemetry *telemetry.Telemetry, emitter *telemetry.Emitter) *Metrics {
	return &Metrics{
		telemetry: telemetry,
		emitter:   emitter,
	}
}

type Metrics struct {
	telemetry *telemetry.Telemetry

	// Telemetry emitter on which to emit metric events.
	emitter *telemetry.Emitter

	// Reusable empty event pool.
	eventPool sync.Pool
}

func (m *Metrics) newEvent() *telemetry.Event {
	event, ok := m.eventPool.Get().(*telemetry.Event)
	if ok {
		return event
	}
	return telemetry.NewEvent(m.telemetry)
}

func (m *Metrics) releaseEvent(event *telemetry.Event) {
	m.eventPool.Put(event)
}

func (m *Metrics) Metric(event *telemetry.Event) {
	m.emitter.Emit(event.WithFields(telemetry.Fields{
		"type": "metric",
	}))
}

func (m *Metrics) Counter(name string, event *telemetry.Event) {
	m.emitter.Emit(event.WithFields(telemetry.Fields{
		"type":        "metric",
		"name":        name,
		"target_type": "counter",
	}))
}

func (m *Metrics) Gauge(name string, event *telemetry.Event) {
	m.emitter.Emit(event.WithFields(telemetry.Fields{
		"type":        "metric",
		"name":        name,
		"target_type": "gauge",
	}))
}

func (m *Metrics) Histogram(name string, event *telemetry.Event) {
	m.emitter.Emit(event.WithFields(telemetry.Fields{
		"type":        "metric",
		"name":        name,
		"target_type": "histogram",
	}))
}

func (m *Metrics) Meter(name string, event *telemetry.Event) {
	m.emitter.Emit(event.WithFields(telemetry.Fields{
		"type":        "metric",
		"name":        name,
		"target_type": "meter",
	}))
}

func (m *Metrics) Timer(name string, event *telemetry.Event) {
	m.emitter.Emit(event.WithFields(telemetry.Fields{
		"type":        "metric",
		"name":        name,
		"target_type": "timer",
	}))
}
