package main

import (
	"time"
)

func main() {
	// Ensure sync on singleton
	for i := 0; i < 10; i++ {
		go getLoggerInstance()
	}
	time.Sleep(1 * time.Millisecond)

	log := getLoggerInstance()

	log.SetLogLevel(1)
	log.Log("This is a log message 1")

	log = getLoggerInstance()
	log.SetLogLevel(2)
	log.Log("This is a log message 2")

	log = getLoggerInstance()
	log.SetLogLevel(3)
	log.Log("This is a log message 3")

	// Wait goroutines
	time.Sleep(50 * time.Millisecond)
}
