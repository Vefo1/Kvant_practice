package logger

import (
	"log"
	"os"
	"strings"
)

// Logger provides a simple logging interface
type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	level       string
}

// NewLogger creates a new Logger instance
func NewLogger(level string) *Logger {
	l := &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		level:       strings.ToLower(level),
	}
	return l
}

// Info logs an informational message
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level == "info" || l.level == "debug" {
		l.infoLogger.Printf(format, v...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.level == "info" || l.level == "warn" || l.level == "debug" {
		l.warnLogger.Printf(format, v...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level == "debug" {
		l.debugLogger.Printf(format, v...)
	}
}

// Fatal logs a fatal error and exits the program
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.errorLogger.Fatalf(format, v...)
}
