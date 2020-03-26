package golog

import "sync"

// Logger struct represents the logger
// name of logger can be set to represent file that is being logged
// or any other specific use.
// path is where the log file will be created
// level represents severity level of logger
// levels available - info, debug, warning, critical
type Logger struct {
	name  string
	path  string
	level Level
	mutex sync.Mutex
}

// Log is a constructor for logger structure
func Log(name string, path string, level Level) *Logger {
	return &Logger{
		name:  name,
		path:  path,
		level: level,
	}
}
