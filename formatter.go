package telemetry

type Formatter interface {
	// Format returns a slice of bytes that will be written to `Writer.Out`.
	Format(map[string]interface{}) ([]byte, error)
}
