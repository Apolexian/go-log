package golog

// Level abstracts enumeration for severity level of logger
type Level uint

// Level represents the specific level of debugger severity
const (
	_          = iota
	INFO Level = iota
	DEBUG
	WARNING
	CRITICAL
)

// levelsMapping maps levels to names of levels
var (
	levelsMapping = map[string]Level{
		"Info":     INFO,
		"Debug":    DEBUG,
		"Warning":  WARNING,
		"Critical": CRITICAL,
	}
)
