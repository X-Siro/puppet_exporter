# puppet_exporter

Simple prometheus exporter. Reads and parse puppet last run report https://puppet.com/docs/puppet/6.18/configuration.html#lastrunreport

This is a partial implementation of a ruby exoprter https://github.com/voxpupuli/puppet-prometheus_reporter/
There is a grafana dashbord for puppet-prometheus-reporter https://grafana.com/grafana/dashboards/700

## Build

`go get && go build .`

## Run
`./puppet_exporter`

## Metrics example
```
# HELP puppet_report Unix timestamp of the last puppet run
# TYPE puppet_report gauge
puppet_report{environment="production",host="zabbix.mgs.local",noop="true",status="unchanged",version="6.17.0"} 1.598550305e+12
# HELP puppet_report_changes Changed resources in the last puppet run
# TYPE puppet_report_changes gauge
puppet_report_changes{environment="production",host="zabbix.mgs.local",name="Total"} 0
# HELP puppet_report_events Resource application events
# TYPE puppet_report_events gauge
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Failure"} 0
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Noop"} 2
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Success"} 0
puppet_report_events{environment="production",host="zabbix.mgs.local",name="Total"} 2
# HELP puppet_report_resources Resources broken down by their state
# TYPE puppet_report_resources gauge
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Changed"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Corrective change"} 1
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Failed"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Failed to restart"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Out of sync"} 1
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Restarted"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Scheduled"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Skipped"} 0
puppet_report_resources{environment="production",host="zabbix.mgs.local",name="Total"} 327
# HELP puppet_report_time Resource apply times
# TYPE puppet_report_time gauge
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Anchor"} 0.0006055450000000001
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Archive"} 0.000223749
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Augeas"} 0.087641602
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Catalog application"} 2.222302386071533
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Concat file"} 0.001062939
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Concat fragment"} 0.0029326000000000005
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Config retrieval"} 2.9746094150468707
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Convert catalog"} 0.8279071138240397
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Exec"} 0.20353354
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Fact generation"} 1.7725586229935288
puppet_report_time{environment="production",host="zabbix.mgs.local",name="File"} 0.09264398099999999
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Filebucket"} 6.2406e-05
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Firewall"} 0.00541201
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Group"} 0.001262543
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Kernel parameter"} 0.006980537
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql database"} 0.000434395
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql datadir"} 0.00039102
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql grant"} 0.000649793
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Mysql user"} 0.000426605
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Node retrieval"} 0.1640859697945416
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Package"} 0.22988559000000008
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Plugin sync"} 0.7470447602681816
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Resources"} 8.7718e-05
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Schedule"} 0.000362534
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Service"} 0.21415389
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Startup time"} 0.658697075
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Total"} 9.385631722
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Transaction evaluation"} 2.1736922720447183
puppet_report_time{environment="production",host="zabbix.mgs.local",name="User"} 0.018924081
puppet_report_time{environment="production",host="zabbix.mgs.local",name="Yumrepo"} 0.001995337
```

