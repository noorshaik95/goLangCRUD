package logger

import (
	"log"
	"os"
)

var GlobalLogger *log.Logger

func init() {
	GlobalLogger = log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(format string, v ...interface{}) {
	GlobalLogger.Printf("INFO: "+format, v...)
}

func Error(format string, v ...interface{}) {
	GlobalLogger.Printf("ERROR: "+format, v...)
}
