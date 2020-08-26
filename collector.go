package main

import (
	"strings"
	"unicode"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector implemantation of prometheus collector
type Collector struct {
	c          int
	config     Config
	chResponse chan *PuppetReport
	chRequest  chan interface{}
}

// Describe func for prometheus collector interface
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	// Just for interface
	ch <- &prometheus.Desc{}
}

// Collect collector method
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	puppetReport := &PuppetReport{}
	collector.chRequest <- struct{}{}
	puppetReport = <-collector.chResponse

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("puppet_report", "Unix timestamp of the last puppet run", []string{"environment", "host", "puppetserver"}, nil),
		prometheus.GaugeValue,
		puppetReport.Version.Config,
		puppetReport.PuppetEnvironment,
		puppetReport.Host,
		puppetReport.PuppetServer,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("puppet_report_changes", "Changed resources in the last puppet run", []string{"name", "environment", "host", "puppetserver"}, nil),
		prometheus.GaugeValue,
		puppetReport.Changes.Total,
		"Total",
		puppetReport.PuppetEnvironment,
		puppetReport.Host,
		puppetReport.PuppetServer,
	)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("puppet_report_events", "Resource application events", []string{"name", "environment", "host", "puppetserver"}, nil),
		prometheus.GaugeValue,
		puppetReport.Events.Failure,
		"Failure",
		puppetReport.PuppetEnvironment,
		puppetReport.Host,
		puppetReport.PuppetServer,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("puppet_report_events", "Resource application events", []string{"name", "environment", "host", "puppetserver"}, nil),
		prometheus.GaugeValue,
		puppetReport.Events.Success,
		"Success",
		puppetReport.PuppetEnvironment,
		puppetReport.Host,
		puppetReport.PuppetServer,
	)
	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc("puppet_report_events", "Resource application events", []string{"name", "environment", "host", "puppetserver"}, nil),
		prometheus.GaugeValue,
		puppetReport.Events.Total,
		"Total",
		puppetReport.PuppetEnvironment,
		puppetReport.Host,
		puppetReport.PuppetServer,
	)

	for key, value := range puppetReport.Time {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("puppet_report_time", "Resource apply times", []string{"name", "environment", "host", "puppetserver"}, nil),
			prometheus.GaugeValue,
			value,
			strings.ReplaceAll(upcaseInitial(key), "_", " "),
			puppetReport.PuppetEnvironment,
			puppetReport.Host,
			puppetReport.PuppetServer,
		)
	}

	for key, value := range puppetReport.Resources {
		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc("puppet_report_resources", "Resources broken down by their state", []string{"name", "environment", "host", "puppetserver"}, nil),
			prometheus.GaugeValue,
			value,
			strings.ReplaceAll(upcaseInitial(key), "_", " "),
			puppetReport.PuppetEnvironment,
			puppetReport.Host,
			puppetReport.PuppetServer,
		)
	}
}

func upcaseInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
