package golog

import (
	"fmt"
	"log"
	"os"
)

// Logger struct represents the logger
// name of logger can be set to represent file that is being logged
// or any other specific use.
// path is where the log file will be created
// level represents severity level of logger
// levels available - info, debug, warning, critical
type Logger struct {
	name     string
	path     string
	Info     *log.Logger
	Debug    *log.Logger
	Warning  *log.Logger
	Critical *log.Logger
}

// Logger flags for file writing
const (
	flags = log.Ldate | log.Lmicroseconds | log.Lshortfile
)

// Initialise logger
func Initialise(name string, path string) *Logger {
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(1)
	}
	return &Logger{
		name:     name,
		path:     path,
		Info:     log.New(logFile, name+": "+levelsMapping[INFO]+":", flags),
		Debug:    log.New(logFile, name+": "+levelsMapping[DEBUG]+":", flags),
		Warning:  log.New(logFile, name+": "+levelsMapping[WARNING]+":", flags),
		Critical: log.New(logFile, name+": "+levelsMapping[CRITICAL]+":", flags),
	}
}

// GetName returns the name of the logger
// e.g name := logger.GetName()
func (log *Logger) GetName() string {
	return log.name
}

// GetPath returns the path of the logger
// e.g path := logger.GetPath()
func (log *Logger) GetPath() string {
	return log.path
}

// SetPath returns a new logger with specified path
func (log *Logger) SetPath(path string) *Logger {
	newLogger := Initialise(log.name, path)
	return newLogger
}

// SetPath returns a new logger with specified name
func (log *Logger) SetName(name string) *Logger {
	newLogger := Initialise(name, log.path)
	return newLogger
}
