package golog

// LogError struct represents errors in module
type LogError struct {
	Err  string `json:"error"`
	Text string `json:"text"`
}

//Error implements the error struct.
func (e LogError) Error() string {
	if e.Err == "" {
		return e.Text
	}
	return e.Err
}
