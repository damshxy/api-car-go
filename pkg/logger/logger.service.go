package logger

import (
	"log"
	"os"
)

type LoggerService interface {
	Info(message string)
	Error(message string)
}

type loggerService struct {
	logger *log.Logger
}

func NewLoggerService() LoggerService {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &loggerService{
		logger: logger,
	}
}

func (l *loggerService) Info(message string) {
	l.logger.Println("INFO: " + message)
}

func (l *loggerService) Error(message string) {
	l.logger.Println("ERROR: " + message)
}