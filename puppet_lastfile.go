package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// PuppetReport - struct for parsing lastrunfile
type PuppetReport struct {
	Version           PuppetReportVersion `yaml:"version"`
	Resources         map[string]float64  `yaml:"resources"`
	Time              map[string]float64  `yaml:"time"`
	Changes           PuppetReportChanges `yaml:"changes"`
	Events            PuppetReportEvents  `yaml:"events"`
	PuppetEnvironment string
	PuppetServer      string
	Host              string
}

// PuppetReportVersion - struct for section 'version' from lastrunfile
type PuppetReportVersion struct {
	Config float64 `yaml:"config"` // epoch of last run start
	Puppet string  `yaml:"puppet"` // version of puppet agent
}

// PuppetReportChanges - struct for section 'total' from lastrunfile
type PuppetReportChanges struct {
	Total float64 `yaml:"total"`
}

// PuppetReportEvents - struct for section 'events' from lastrunfile
type PuppetReportEvents struct {
	Failure float64 `yaml:"failure"`
	Success float64 `yaml:"success"`
	Total   float64 `yaml:"total"`
}

// ReadFile lastrunfile to PuppetReport object
func (report *PuppetReport) ReadFile(lastrunfile string) (err error) {
	reportFile, err := ioutil.ReadFile(lastrunfile)
	if err != nil {
		log.Printf("ioutil.ReadFile error: #%v ", err)
	}
	err = yaml.Unmarshal(reportFile, report)
	if err != nil {
		log.Fatalf("Unmarshal lastrunfile error: %v", err)
	}
	return nil
}
