package telemetry

var ErrorKey = "error"
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

// Event implements the Error interface so that structured telemetry can be returned everywhere an error can.
func (event Event) Error() string {
	err, exists := event.data[ErrorKey]
	if !exists {
		return toString(event)
	}
	str, ok := err.(string)
	if !ok {
		return toString(event)
	}
	return str
}

// Default serializer for use with Error interface implementation.
func toString(event Event) string {
	formatter := &JSONFormatter{}
	str, err := formatter.Format(event.Marshal())
	if err != nil {
		return err.Error()
	}
	return string(str)
}
