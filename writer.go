package telemetry

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

var bufferPool *sync.Pool

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// Creates new Writer instance. You can set where events are written by changing the `Out` property.
// You can change the format of what is written by changing the `Formatter` property.
func NewWriter() *Writer {
	return &Writer{
		Out:       os.Stdout,
		Formatter: new(JSONFormatter),
	}
}

type Writer struct {
	// The writer writes out via `io.Copy` to this in a mutex. Default is `os.Stdout`, but can be set to any
	// `io.Writer`.
	Out io.Writer

	// All writes pass through the formatter before being written to Out. `JSONFormatter` is the default.
	Formatter Formatter

	// Use for locking when writing to Out.
	mutex sync.Mutex
}

func (writer Writer) Write(event map[string]interface{}) {
	var buffer *bytes.Buffer
	buffer = bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)
	serialized, err := writer.Formatter.Format(event)
	if err != nil {
		message := fmt.Sprintf("{\"type\":\"log\",\"level\":\"error\",\"message\":\"%v\"}\n", err)
		writer.mutex.Lock()
		_, err = writer.Out.Write([]byte(message))
		if err != nil {
			fmt.Fprintf(os.Stderr, message)
		}
		writer.mutex.Unlock()
		return
	}
	writer.mutex.Lock()
	_, err = writer.Out.Write(serialized)
	if err != nil {
		fmt.Fprintf(os.Stderr, "{\"type\":\"log\",\"level\":\"error\",\"message\":\"%v\"}\n", err)
	}
	writer.mutex.Unlock()
}
