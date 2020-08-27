package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector implemantation of prometheus collector
type Collector struct {
	c          int
	config     Config
	chResponse chan *PuppetDetailedReport
	chRequest  chan interface{}
}

// Describe func for prometheus collector interface
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	// Just for interface
	ch <- &prometheus.Desc{}
}

// Collect collector method
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	puppetDetailedReport := &PuppetDetailedReport{}
	collector.chRequest <- struct{}{}
	puppetDetailedReport = <-collector.chResponse

	// ch <- prometheus.MustNewConstMetric(
	// 	prometheus.NewDesc("puppet_report", "Unix timestamp of the last puppet run", []string{"environment", "host", "puppetserver"}, nil),
	// 	prometheus.GaugeValue,
	// 	puppetSummaryReport.Version.Config,
	// 	puppetSummaryReport.PuppetEnvironment,
	// 	puppetSummaryReport.Host,
	// 	puppetSummaryReport.PuppetServer,
	// )
	fmt.Println("Host: ", puppetDetailedReport.Host)
	fmt.Println("Time: ", puppetDetailedReport.Time)
	fmt.Println("ConfigurationVersion: ", puppetDetailedReport.ConfigurationVersion)
	fmt.Println("TransactionUUID: ", puppetDetailedReport.TransactionUUID)
	fmt.Println("ReportFormat: ", puppetDetailedReport.ReportFormat)
	fmt.Println("PuppetVersion: ", puppetDetailedReport.PuppetVersion)
	fmt.Println("Status: ", puppetDetailedReport.Status)
	fmt.Println("Noop: ", puppetDetailedReport.Noop)
	fmt.Println("NoopPending: ", puppetDetailedReport.NoopPending)
	fmt.Println("Environment: ", puppetDetailedReport.Environment)
	fmt.Println("CorrectiveChange: ", puppetDetailedReport.CorrectiveChange)
	fmt.Println("CatalogUUID: ", puppetDetailedReport.CatalogUUID)
	fmt.Println("CachedCatalogStatus: ", puppetDetailedReport.CachedCatalogStatus)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("puppet_report", "Unix timestamp of the last puppet run", []string{"environment", "host", "version", "noop", "status"}, nil),
		prometheus.GaugeValue,
		puppetDetailedReport.ConfigurationVersion*1000, // for dashboard compability
		puppetDetailedReport.Environment,
		puppetDetailedReport.Host,
		puppetDetailedReport.PuppetVersion,
		fmt.Sprintf("%v", puppetDetailedReport.Noop),
		puppetDetailedReport.Status,
	)

	for name, metric := range puppetDetailedReport.Metrics {
		description := ""
		switch name {
		case "resources":
			description = "Resources broken down by their state"
		case "time":
			description = "Resource apply times"
		case "events":
			description = "Resource application events"
		case "changes":
			description = "Changed resources in the last puppet run"
		}
		for _, metricValue := range metric.Values {
			val, err := strconv.ParseFloat(metricValue[2], 64)
			if err != nil {
				log.Fatal(err)
			}
			ch <- prometheus.MustNewConstMetric(
				prometheus.NewDesc("puppet_report_"+name, description, []string{"name", "environment", "host"}, nil),
				prometheus.GaugeValue,
				val,
				metricValue[1],
				puppetDetailedReport.Environment,
				puppetDetailedReport.Host,
			)
		}
	}
}
