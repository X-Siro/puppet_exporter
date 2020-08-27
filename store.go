package main

import (
	"log"
	"time"
)

// Store - goroutine for geting puppet last run status with caching
//   chRequest  - channel for get request for puppet last run status
//   chResponse - channels for response puppet last run status
func Store(config Config, chRequest chan interface{}, chResponse chan *PuppetDetailedReport) {
	timestamp := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	puppetDetailedReport := &PuppetDetailedReport{}

	for range chRequest {
		if time.Now().Sub(timestamp) > config.lastRunFileCache {
			err := puppetDetailedReport.ReadFile(config.lastRunReport)
			if err != nil {
				log.Fatal(err)
			}
			timestamp = time.Now()
		}

		chResponse <- puppetDetailedReport
	}
}
