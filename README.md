# puppet_exporter

Simple prometheus exporter. Reads and parse puppet last run file https://puppet.com/docs/puppet/6.18/configuration.html#lastrunfile
Labels `puppetserver` and `environment` takes from puppet.conf.

This is a partial implementation of a ruby exoprter https://github.com/voxpupuli/puppet-prometheus_reporter/

## Build

Run `go get && go build .`

## Run
`./prometheus_exporter`

## Metrics example
```
# HELP puppet_report Unix timestamp of the last puppet run
# TYPE puppet_report gauge
puppet_report{environment="production",host="zabbix.mgs.local",puppetserver="puppet.mgs.local"} 1.598452381e+09
# HELP puppet_report_changes Changed resources in the last puppet run
# TYPE puppet_report_changes gauge
puppet_report_changes{environment="production",host="zabbix.mgs.local",name="Total",puppetserver="puppet.mgs.local"} 54
# HELP puppet_report_events Resource application events
# TYPE puppet_report_events gauge
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Failure",puppetserver="puppet.mgs.local"} 0
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Success",puppetserver="puppet.mgs.local"} 54
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Total",puppetserver="puppet.mgs.local"} 54
# HELP puppet_report_resources Resources broken down by their state
# TYPE puppet_report_resources gauge
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Changed",puppetserver="puppet.mgs.local"} 48
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Corrective change",puppetserver="puppet.mgs.local"} 14
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Failed",puppetserver="puppet.mgs.local"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Failed to restart",puppetserver="puppet.mgs.local"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Out of sync",puppetserver="puppet.mgs.local"} 48
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Restarted",puppetserver="puppet.mgs.local"} 4
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Scheduled",puppetserver="puppet.mgs.local"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Skipped",puppetserver="puppet.mgs.local"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Total",puppetserver="puppet.mgs.local"} 344
# HELP puppet_report_time Resource apply times
# TYPE puppet_report_time gauge
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Anchor",puppetserver="puppet.mgs.local"} 0.00037241999999999993
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Archive",puppetserver="puppet.mgs.local"} 115.389414038
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Augeas",puppetserver="puppet.mgs.local"} 0.111242472
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Catalog application",puppetserver="puppet.mgs.local"} 134.70196412783116
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Concat file",puppetserver="puppet.mgs.local"} 0.0012272290000000003
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Concat fragment",puppetserver="puppet.mgs.local"} 0.005865848
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Config retrieval",puppetserver="puppet.mgs.local"} 2.3082161438651383
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Convert catalog",puppetserver="puppet.mgs.local"} 0.8946966370567679
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Exec",puppetserver="puppet.mgs.local"} 0.22742444999999997
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Fact generation",puppetserver="puppet.mgs.local"} 1.8536186451092362
puppet_report_time{environment="production",host="zabbix.mgs.local",name="File",puppetserver="puppet.mgs.local"} 0.36711469699999993
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Filebucket",puppetserver="puppet.mgs.local"} 6.1669e-05
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Firewall",puppetserver="puppet.mgs.local"} 0.23260996100000003
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Group",puppetserver="puppet.mgs.local"} 0.037651976999999996
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Kernel parameter",puppetserver="puppet.mgs.local"} 0.007604699
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Last run",puppetserver="puppet.mgs.local"} 1.598452519e+09
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql database",puppetserver="puppet.mgs.local"} 0.000860453
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql datadir",puppetserver="puppet.mgs.local"} 0.000620924
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql grant",puppetserver="puppet.mgs.local"} 0.000714767
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql user",puppetserver="puppet.mgs.local"} 0.000582597
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Node retrieval",puppetserver="puppet.mgs.local"} 0.18439693469554186
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Package",puppetserver="puppet.mgs.local"} 14.290790266
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Plugin sync",puppetserver="puppet.mgs.local"} 0.7674825801514089
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Resources",puppetserver="puppet.mgs.local"} 5.0147e-05
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Schedule",puppetserver="puppet.mgs.local"} 0.000394482
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Service",puppetserver="puppet.mgs.local"} 0.767115533
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Startup time",puppetserver="puppet.mgs.local"} 0.699846496
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Total",puppetserver="puppet.mgs.local"} 141.585556573
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Transaction evaluation",puppetserver="puppet.mgs.local"} 134.675112857949
puppet_report_time{environment="production",host="zabbix.mgs.local",name="User",puppetserver="puppet.mgs.local"} 0.06356155000000001
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Yumrepo",puppetserver="puppet.mgs.local"} 0.01740214
```
