package main

type LogFiles map[string]interface{}

type LoggerAdapter interface {
	Info(msg string, fileds LogFiles)
	Debug(msg string, fileds LogFiles)
	Error(msg string, err error, fileds LogFiles)
	With(fieds LogFiles) LoggerAdapter
}
