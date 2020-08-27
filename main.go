package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	arguments, err := docopt.Parse(usage, nil, true, "v1.0.1", true, false)
	if err != nil {
		fmt.Printf("\nRun for print help screen:\n\t%v --help\n", os.Args[0])
		os.Exit(1)
	}

	// Exit if setted help or version flags
	if os.Args != nil && len(arguments) == 0 {
		os.Exit(0)
	}

	config := Config{}
	err = config.ReadArgs(arguments)
	if err != nil {
		log.Fatal(err)
	}

	chResponse := make(chan *PuppetDetailedReport, 10)
	chRequest := make(chan interface{}, 10)
	defer close(chRequest)
	defer close(chResponse)
	go Store(config, chRequest, chResponse)

	collector := &Collector{
		c:          0,
		config:     config,
		chResponse: chResponse,
		chRequest:  chRequest,
	}

	prometheus.MustRegister(collector)
	http.Handle("/", promhttp.Handler())
	fmt.Printf("puppet_exporter listen on %v\n", config.listen)
	if err := http.ListenAndServe(config.listen, nil); err != nil {
		log.Fatal(err)
	}
}
