package logger

import (
	"fmt"
	"sync"

	"github.com/tristanls/telemetry"
	"strings"
)

// Log level type
type Level uint8

const (
	Fatal Level = iota
	Error
	Warn
	Info
	Debug
)

func (level Level) String() string {
	switch level {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	}

	return "unknown"
}

func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return Debug, nil
	case "info":
		return Info, nil
	case "warn":
		return Warn, nil
	case "error":
		return Error, nil
	case "fatal":
		return Fatal, nil
	}

	var l Level
	return l, telemetry.New().WithFields(telemetry.Fields{
		"error": "Expected log level to be one of: debug, info, warn, error, fatal.",
		"args":  []string{level},
	})
}

// Creates new Logger instance that will emit Events on provided telemetry emitter using provided telemetry
// configuration.
func NewLogger(telemetry *telemetry.Telemetry, emitter *telemetry.Emitter) *Logger {
	return &Logger{
		telemetry: telemetry,
		emitter:   emitter,
		level:     Info,
	}
}

func NewLoggerWithLevel(telemetry *telemetry.Telemetry, emitter *telemetry.Emitter, level Level) *Logger {
	return &Logger{
		telemetry: telemetry,
		emitter:   emitter,
		level:     level,
	}
}

type Logger struct {
	telemetry *telemetry.Telemetry

	// Telemetry emitter on which to emit log events.
	emitter *telemetry.Emitter

	// Logger's log level
	level Level

	// Reusable empty event pool.
	eventPool sync.Pool
}

func (logger *Logger) newEvent() *telemetry.Event {
	event, ok := logger.eventPool.Get().(*telemetry.Event)
	if ok {
		return event
	}
	return telemetry.NewEvent(logger.telemetry)
}

func (logger *Logger) releaseEvent(event *telemetry.Event) {
	logger.eventPool.Put(event)
}

func (logger *Logger) Log(level Level, args ...interface{}) {
	if logger.level >= level {
		event := logger.newEvent()
		defer logger.releaseEvent(event)
		if len(args) > 0 {
			logger.emitter.Emit(event.WithFields(telemetry.Fields{
				"type":    "log",
				"level":   level.String(),
				"message": fmt.Sprint(args...),
			}))
		} else {
			logger.emitter.Emit(event.WithFields(telemetry.Fields{
				"type":  "log",
				"level": level.String(),
			}))
		}
	}
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	if logger.level >= level {
		event := logger.newEvent()
		defer logger.releaseEvent(event)
		logger.emitter.Emit(event.WithFields(telemetry.Fields{
			"type":    "log",
			"level":   level.String(),
			"message": fmt.Sprintf(format, args...),
		}))
	}
}

func (logger *Logger) Loge(level Level, event *telemetry.Event, args ...interface{}) {
	if logger.level >= level {
		if len(args) > 0 {
			logger.emitter.Emit(event.WithFields(telemetry.Fields{
				"type":    "log",
				"level":   level.String(),
				"message": fmt.Sprint(args...),
			}))
		} else {
			logger.emitter.Emit(event.WithFields(telemetry.Fields{
				"type":  "log",
				"level": level.String(),
			}))
		}
	}
}

func (logger *Logger) Logef(level Level, event *telemetry.Event, format string, args ...interface{}) {
	if logger.level >= level {
		logger.emitter.Emit(event.WithFields(telemetry.Fields{
			"type":    "log",
			"level":   level.String(),
			"message": fmt.Sprintf(format, args...),
		}))
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Log(Debug, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Logf(Debug, format, args...)
}

func (logger *Logger) Debuge(event *telemetry.Event, args ...interface{}) {
	logger.Loge(Debug, event, args...)
}

func (logger *Logger) Debugef(event *telemetry.Event, format string, args ...interface{}) {
	logger.Logef(Debug, event, format, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Log(Info, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Logf(Info, format, args...)
}

func (logger *Logger) Infoe(event *telemetry.Event, args ...interface{}) {
	logger.Loge(Info, event, args...)
}

func (logger *Logger) Infoef(event *telemetry.Event, format string, args ...interface{}) {
	logger.Logef(Info, event, format, args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.Log(Warn, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Logf(Warn, format, args...)
}

func (logger *Logger) Warne(event *telemetry.Event, args ...interface{}) {
	logger.Loge(Warn, event, args...)
}

func (logger *Logger) Warnef(event *telemetry.Event, format string, args ...interface{}) {
	logger.Logef(Warn, event, format, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Log(Error, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Logf(Error, format, args...)
}

func (logger *Logger) Errore(event *telemetry.Event, args ...interface{}) {
	logger.Loge(Error, event, args...)
}

func (logger *Logger) Erroref(event *telemetry.Event, format string, args ...interface{}) {
	logger.Logef(Error, event, format, args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.Log(Fatal, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.Logf(Fatal, format, args...)
}

func (logger *Logger) Fatale(event *telemetry.Event, args ...interface{}) {
	logger.Loge(Fatal, event, args...)
}

func (logger *Logger) Fatalef(event *telemetry.Event, format string, args ...interface{}) {
	logger.Logef(Fatal, event, format, args...)
}
