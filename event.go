package telemetry

var ProvenanceKey = "provenance"
var TimestampKey = "timestamp"

func NewEvent(telemetry *Telemetry) *Event {
	return &Event{
		telemetry: telemetry,
		data:      make(Fields, 5),
	}
}

// Event maintains track of event data fields and provenance of the event.
// It provides affordances for adding fields and provenance as needed.
type Event struct {
	telemetry *Telemetry

	// Contains a list (slice) of provenance of the event.
	provenance []Fields

	// Contains all non-provenance fields.
	data Fields
}

func (event *Event) WithField(key string, value interface{}) *Event {
	return event.WithFields(map[string]interface{}{key: value})
}

// Adds specified fields to the Event and returns new Event with those fields included.
func (event *Event) WithFields(fields Fields) *Event {
	data := make(Fields, len(event.data)+len(fields))
	for k, v := range event.data {
		data[k] = v
	}
	for k, v := range fields {
		data[k] = v
	}
	return &Event{telemetry: event.telemetry, provenance: event.provenance, data: data}
}

// Adds specified provenance to the Event and returns new Event with that provenance included.
func (event *Event) WithProvenance(fields Fields) *Event {
	var provenance []Fields
	if event.provenance == nil {
		provenance = make([]Fields, 1)
		provenance[0] = fields
	} else {
		provenance = make([]Fields, len(event.provenance)+1)
		for i, p := range event.provenance {
			provenance[i] = p
		}
		provenance[len(provenance)-1] = fields
	}
	return &Event{telemetry: event.telemetry, provenance: provenance, data: event.data}
}

// Convert Event into a map[string]interface{} for logging, transport, or other such usage once no further
// enrichment of the Event is needed.
func (event *Event) Marshal() map[string]interface{} {
	fields := make(map[string]interface{}, len(event.data)+1)
	if event.provenance != nil {
		fields[ProvenanceKey] = event.provenance
	}
	for k, v := range event.data {
		fields[k] = v
	}
	return fields
}
