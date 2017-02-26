package telemetry

func Join(base, more Fields) Fields {
	data := make(Fields, len(base)+len(more))
	for k, v := range base {
		data[k] = v
	}
	for k, v := range more {
		if existing, exists := data[k]; exists {
			if _, ok := existing.(Fields); ok {
				if _, ok := v.(Fields); ok {
					data[k] = Join(existing.(Fields), v.(Fields))
				} else {
					data[k] = v // override if not joining Fields
				}
			} else {
				data[k] = v // override if not joining Fields
			}
		} else {
			data[k] = v
		}
	}
	return data
}