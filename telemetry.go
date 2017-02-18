package telemetry

import (
	"sync"
)

type Fields map[string]interface{}

// Creates new Telemetry instance. You can set what layout is used for timestamp formatting by
// changing `TimestampLayout` property.
func New() *Telemetry {
	return &Telemetry{
		TimestampLayout: "2006-01-02T15:04:05.000Z",
	}
}

type Telemetry struct {
	// Timestamp layout to be used with time.Format(layout) when creating timestamp for an event.
	// Default is "2006-01-02T15:04:05.000Z". Timestamps are always in UTC timezone (Z), so make
	// sure your layout includes timezone information or assumes UTC.
	TimestampLayout string

	// Reusable empty event pool.
	eventPool sync.Pool
}

func (t *Telemetry) newEvent() *Event {
	event, ok := t.eventPool.Get().(*Event)
	if ok {
		return event
	}
	return NewEvent(t)
}

func (t *Telemetry) releaseEvent(event *Event) {
	t.eventPool.Put(event)
}

func (t *Telemetry) WithField(key string, value interface{}) *Event {
	event := t.newEvent()
	defer t.releaseEvent(event)
	return event.WithField(key, value)
}

func (t *Telemetry) WithFields(fields Fields) *Event {
	event := t.newEvent()
	defer t.releaseEvent(event)
	return event.WithFields(fields)
}

func (t *Telemetry) WithProvenance(fields Fields) *Event {
	event := t.newEvent()
	defer t.releaseEvent(event)
	return event.WithProvenance(fields)
}
