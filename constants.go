package main

const usage = `Usage:
  puppet_exporter [options]

Options:
  -h, --help                        Show this screen.
  -l, --listen=[<host>]:<port>		  IP:port or ':port' for listening on 0.0.0.0. Default is ':9101'
  -t, --lastrunfilecache=<seconds>	TTL of cache puppet lastrunfile data. Default is 600 seconds
  -r, --lastrunreport=<file>        File with puppet last run report. Default is /opt/puppetlabs/puppet/cache/state/last_run_report.yaml
  -c, --puppetconf=<file>			      Puppet conf file. Default is /etc/puppetlabs/puppet/puppet.conf
`
const (
	defaultListen           = ":9101"
	defaultLastRunFileCache = 600
	defaultLastRunReport    = "/opt/puppetlabs/puppet/cache/state/last_run_report.yaml"
	defaultPuppetConf       = "/etc/puppetlabs/puppet/puppet.conf"
)
