package main

import (
	"log"
	"os"
	//"fmt"
)

// aggerate log
type aggregatedLog struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func newAggregatedLog() *aggregatedLog {
	return &aggregatedLog{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func main() {

	log.Println("This is a log message")

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("This is a log message")

	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("This is a log message")

	log.Panic("This is a panic")
	log.Fatalln("This is a fatal error")

	aggregatedLog := newAggregatedLog()

	aggregatedLog.infoLogger.Println("This is a info log message")
	aggregatedLog.warningLogger.Println("This is a warning log message")
	aggregatedLog.errorLogger.Println("This is a error log message")

}
