package telemetry

import (
	"encoding/json"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(event map[string]interface{}) ([]byte, error) {
	data := make(map[string]interface{}, len(event))
	for k, v := range event {
		switch v := v.(type) {
		case error:
			// Copied from logrus implementation because errors are ignored by `encoding/json`
			// as explained in https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return append(serialized, '\n'), nil
}
