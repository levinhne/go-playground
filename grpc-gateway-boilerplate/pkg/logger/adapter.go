package logger

type LogFields map[string]interface{}

type LoggerAdapter interface {
	Info(msg string, fileds LogFields)
	Debug(msg string, fileds LogFields)
	Error(msg string, err error, fileds LogFields)
	With(fieds LogFields) LoggerAdapter
}

// NopLogger is a logger which discards all logs.
type NopLogger struct{}

func (NopLogger) Error(msg string, err error, fields LogFields) {}
func (NopLogger) Info(msg string, fields LogFields)             {}
func (NopLogger) Debug(msg string, fields LogFields)            {}
func (NopLogger) Trace(msg string, fields LogFields)            {}
func (l NopLogger) With(fields LogFields) LoggerAdapter         { return l }
