package main

import "../golog"

func main() {
	// create a new logger instance with:
	// name: myLogger
	// path: myLog.log (path counts from project basedir)
	myLogger := golog.Initialise("myLogger","myLog.log")
	// output to logging file with level Info
	myLogger.Info.Print("Logger was created!")
	// to change name of myLogger use SetName and create new instance of logger
	myLogger = myLogger.SetName("myNewLogger")
	// to change path of myLogger use SetPath and create new instance of logger
	myLogger = myLogger.SetPath("myNewLog.log")
	myLogger.Info.Print("This is now the new logger!")
}
