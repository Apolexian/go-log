package main

import (
	"../golog"
)

func main() {
	logger := golog.Initialise("test", "test.log")
	logger.Warning.Print("This is a test")
	logger = logger.SetPath("test2.log")
	logger.Debug.Print("This is a test")
	logger = logger.SetName("new test")
	logger.Info.Print("This is a test")
}
