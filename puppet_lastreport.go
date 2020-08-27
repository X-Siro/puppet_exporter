package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// PuppetDetailedReport - struct for parsing lastrun report file
type PuppetDetailedReport struct {
	Host                 string                                        `yaml:"host"`
	Time                 string                                        `yaml:"time"`
	ConfigurationVersion float64                                       `yaml:"configuration_version"`
	TransactionUUID      string                                        `yaml:"transaction_uuid"`
	ReportFormat         float64                                       `yaml:"report_format"`
	PuppetVersion        string                                        `yaml:"puppet_version"`
	Status               string                                        `yaml:"status"`
	TransactionCompleted bool                                          `yaml:"transaction_completed"`
	Noop                 bool                                          `yaml:"noop"`
	NoopPending          bool                                          `yaml:"bool_pending"`
	Environment          string                                        `yaml:"environment"`
	Logs                 []PuppetDetailedReportLog                     `yaml:"logs"`
	CorrectiveChange     bool                                          `yaml:"corrective_change"`
	CatalogUUID          string                                        `yaml:"catalog_uuid"`
	CachedCatalogStatus  string                                        `yaml:"cached_catalog_status"`
	ResourceStatuses     map[string]PuppetDetailedReportResourceStatus `yaml:"resource_statuses"`
	Metrics              map[string]PuppetDetailedReportMetric         `yaml:"metrics"`
}

// PuppetDetailedReportLog - struct for log object in lastrun report file
type PuppetDetailedReportLog struct {
	Level   string   `yaml:"level"`
	Message string   `yaml:"message"`
	Source  string   `yaml:"source"`
	Tags    []string `yaml:"tags"`
	Time    string   `yaml:"time"`
	Line    int      `yaml:"line"`
}

// PuppetDetailedReportMetric -
type PuppetDetailedReportMetric struct {
	Name   string     `yaml:"name"`
	Label  string     `yaml:"label"`
	Values [][]string `yaml:"values"`
}

// // PuppetDetailedReportMetricsValue -
// type PuppetDetailedReportMetricsValue struct {
// 	Name  string
// 	Label string
// 	Value float64
// }

// PuppetDetailedReportResourceStatus -
type PuppetDetailedReportResourceStatus struct {
	Title            string                                  `yaml:"title"`
	File             string                                  `yaml:"file"`
	Line             int                                     `yaml:"line"`
	Resource         string                                  `yaml:"resource"`
	ResourceType     string                                  `yaml:"resource_type"`
	ProviderUsed     string                                  `yaml:"provider_used"`
	ContainmentPath  []string                                `yaml:"containment_path"`
	EvaluationTime   float64                                 `yaml:"evaluation_time"`
	Tags             []string                                `yaml:"tags"`
	Time             string                                  `yaml:"time"`
	Failed           bool                                    `yaml:"failed"`
	FailedToRestart  bool                                    `yaml:"failed_to_restart"`
	Changed          bool                                    `yaml:"changed"`
	OutOfSync        bool                                    `yaml:"out_of_sync"`
	Skipped          bool                                    `yaml:"skipped"`
	ChangeCount      int                                     `yaml:"change_count"`
	Events           []PuppetDetailedReportResourceStatusLog `yaml:"events"`
	CorrectiveChange bool                                    `yaml:"corrective_change"`
}

// PuppetDetailedReportResourceStatusLog -
type PuppetDetailedReportResourceStatusLog struct {
	Property         string `yaml:"property"`
	HistoricalValue  string `yaml:"historical_value"`
	Message          string `yaml:"message"`
	Name             string `yaml:"name"`
	Status           string `yaml:"status"`
	Time             string `yaml:"time"`
	CorrectiveChange bool   `yaml:"corrective_change"`
}

// ReadFile lastrunfile to PuppetReport object
func (report *PuppetDetailedReport) ReadFile(lastrunreport string) (err error) {
	reportFile, err := ioutil.ReadFile(lastrunreport)
	if err != nil {
		log.Printf("ioutil.ReadFile error: #%v ", err)
	}
	err = yaml.Unmarshal(reportFile, report)
	if err != nil {
		log.Fatalf("Unmarshal lastrunfile error: %v", err)
	}
	return nil
}
