package golog

// Level abstracts enumeration for severity level of logger
type Level uint

// Level represents the specific level of debugger severity
const (
	_               = iota
	LevelInfo Level = iota
	LevelDebug
	LevelWarning
	LevelCritical
)
