package data

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ltruelove/gohome/config"
)

var logFile *os.File

func LogFileTicker(Config config.Configuration) {
	ticker := time.NewTicker(24 * time.Hour)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("Logger tick at ", t)

			logFile.Close()

			currentTime := time.Now()
			datedLogFile := fmt.Sprintf("%s_%s", currentTime.Format("2006-01-02"), Config.LogFile)
			logFile, logErr := os.OpenFile(datedLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

			if logErr != nil {
				log.Fatalf("error opening file: %v", logErr)
			}

			defer logFile.Close()
			log.SetOutput(logFile)
		}
	}
}
