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

// SetName is used to set the name of the logger
// e.g logger.SetName("Testing")
func (log *Logger) SetName(name string) {
	log.name = name
}

// GetName returns the name of the logger
// e.g name := logger.GetName()
func (log *Logger) GetName() string {
	return log.name
}

// SetLevel is used to set the level of the logger
// e.g logger.SetLevel(INFO)
func (log *Logger) SetLevel(level Level) {
	log.level = level
}

// GetLevel returns the level of the logger
// e.g level := logger.GetLevel()
func (log *Logger) GetLevel() Level {
	return log.level
}

// SetPath is used to set the path of the logger
// e.g logger.SetPath(path) where path is string
// representing the path
func (log *Logger) SetPath(path string) {
	log.path = path
}

// GetPath returns the path of the logger
// e.g path := logger.GetPath()
func (log *Logger) GetPath() string {
	return log.path
}
