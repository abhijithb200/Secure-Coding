package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // used as a switch to send close singal

func main() {
	go logger()
	logCh <- logEntry{logInfo, "App is starting"}
	logCh <- logEntry{logInfo, "App is shutting down"}

	time.Sleep(100 * time.Millisecond)

	doneCh <- struct{}{} // passing an empty struct signal to doneCh
}

func logger() {
	for {
		select {
		case entry := <-logCh: // open when theere is anyting from logCh
			fmt.Printf("%v - %v\n", entry.severity, entry.message)
		case <-doneCh: // open when there is some singnal in doneCh
			break
		}
	}
}
