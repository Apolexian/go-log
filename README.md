# Go log
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Apolexian/go-log/blob/master/LICENSE)

A wrapper for the go log library aimed to provide a fast way to setup a logger with a nicer format.

## Setup and usage

```go
package example

import "github.com/Apolexian/go-log/golog"

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

```

## Testing
Tests provided in test directory. Run:

```bash
go test -v
```