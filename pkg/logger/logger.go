package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func New() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string, keyvals ...interface{}) {
	l.infoLogger.Println(append([]interface{}{msg}, keyvals...)...)
}

func (l *Logger) Error(msg string, keyvals ...interface{}) {
	l.errorLogger.Println(append([]interface{}{msg}, keyvals...)...)
}
