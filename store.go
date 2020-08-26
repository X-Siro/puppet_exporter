package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

// Store - goroutine for geting puppet last run status with caching
//   chRequest  - channel for get request for puppet last run status
//   chResponse - channels for response puppet last run status
func Store(config Config, chRequest chan interface{}, chResponse chan *PuppetReport) {
	timestamp := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	puppetReport := &PuppetReport{}

	for range chRequest {
		if time.Now().Sub(timestamp) > config.lastRunFileCache {
			err := puppetReport.ReadFile(config.lastRunFile)
			if err != nil {
				log.Fatal(err)
			}
			timestamp = time.Now()
		}

		puppetConf, err := ini.Load(config.puppetConf)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
		}

		puppetReport.PuppetEnvironment = puppetConf.Section("agent").Key("environment").String()
		puppetReport.PuppetServer = puppetConf.Section("agent").Key("server").String()
		puppetReport.Host, err = os.Hostname()
		if err != nil {
			log.Fatalf("Cannot get os.Hostname: %v\n", err)
		}

		chResponse <- puppetReport
	}
}
