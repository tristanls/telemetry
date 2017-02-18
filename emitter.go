package telemetry

import (
	"sync"
	"time"
)

type Listener func(*Event)

func NewEmitter() *Emitter {
	return &Emitter{}
}

type Emitter struct {
	// Collection of Listeners listening for Telemetry Events
	listeners []Listener

	mutex sync.Mutex
}

func (e *Emitter) AddListener(listener Listener) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.listeners = append(e.listeners, listener)
}

func (e *Emitter) RemoveListener(listener Listener) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for i, l := range e.listeners {
		if &l == &listener {
			e.listeners = append(e.listeners[:i], e.listeners[i+1:]...)
			return
		}
	}
}

// Emits the event, adding timestamp if not already present.
func (e *Emitter) Emit(event *Event) {
	var ev *Event
	_, exists := event.data[TimestampKey]
	if !exists {
		ev = event.WithField(TimestampKey, time.Now().UTC().Format(event.telemetry.TimestampLayout))
	} else {
		ev = event.WithFields(nil) // make a copy
	}
	for _, handle := range e.listeners {
		handle(ev)
	}
}
