package telemetry

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJoinFieldsWithFields(t *testing.T) {
	base := Fields{
		"key1": "value1",
		"key2": Fields{
			"key2.1": "value2.1",
		},
		"key3": Fields{
			"key3.1": Fields {
				"key3.1.1": "value3.1.1",
			},
		},
		"key4": Fields{
			"key4.1": "value4.1",
		},
	}
	more := Fields{
		"key1": "value2",
		"key2": Fields{
			"key2.2": "value2.2",
		},
		"key3": Fields{
			"key3.1": Fields{
				"key3.1.2": "value3.1.2",
			},
		},
		"key4": "type_mismatch_override",
	}
	result := Join(base, more)
	require.Equal(t, "value2", result["key1"])
	require.Equal(t, "value2.1", result["key2"].(Fields)["key2.1"])
	require.Equal(t, "value2.2", result["key2"].(Fields)["key2.2"])
	require.Equal(t, "value3.1.1", result["key3"].(Fields)["key3.1"].(Fields)["key3.1.1"])
	require.Equal(t, "value3.1.2", result["key3"].(Fields)["key3.1"].(Fields)["key3.1.2"])
	require.Equal(t, "type_mismatch_override", result["key4"])
}
