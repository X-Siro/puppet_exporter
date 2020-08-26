package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Config             - struct of puppet_export configuration
//   LastRunFile      - path to puppet lastrunfile (https://puppet.com/docs/puppet/6.18/configuration.html#lastrunfile)
//   LastRunFileCache - puppet_exporter reads LastRunFile with this interval
//   Listen           - puppet_exporter listens this host:port
//   PuppetConf		  - path to puppet.conf file whith agent section
type Config struct {
	lastRunFile      string
	lastRunFileCache time.Duration
	listen           string
	puppetConf       string
}

// ReadArgs from command line flags to config object
func (c *Config) ReadArgs(args map[string]interface{}) (err error) {
	c.lastRunFile = defaultLastRunFile
	c.lastRunFileCache = defaultLastRunFileCache
	c.listen = defaultListen
	c.puppetConf = defaultPuppetConf

	// LastRunFile
	if args["--lastrunfile"] != nil {
		c.lastRunFile = args["--lastrunfile"].(string)
	}
	// LastRunFileCache
	if args["--lastrunfilecache"] != nil {
		seconds, err := strconv.Atoi(args["--lastrunfilecache"].(string))
		if err != nil {
			return fmt.Errorf("--lastrunfilecache wrong seconds argument: '%v' is not numeric", args["--lastrunfilecache"].(string))
		}
		c.lastRunFileCache = time.Duration(seconds) * time.Second
	}

	//Listen
	if args["--listen"] != nil {
		listen := strings.Split(args["--listen"].(string), ":")
		// check if arg not have ':'
		if len(listen) < 2 {
			return fmt.Errorf("--listen wrong format")
		}
		// check if port is not numeric
		if _, err = strconv.Atoi(listen[1]); err != nil {
			return fmt.Errorf("--listen wrong argument: port '%v' is not numeric", listen[1])
		}

		if listen[0] == "" {
			c.listen = fmt.Sprintf(":%v", listen[1])
		}
	}

	// puppetConf
	if args["--puppetconf"] != nil {
		c.puppetConf = args["--puppetconf"].(string)
	}

	return nil
}
