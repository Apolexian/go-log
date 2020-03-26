package main

import (
	"../golog"
)

func main() {
	logger := golog.Initialise("test", "test.log")
	logger.Warning.Print("This is a test")
}
