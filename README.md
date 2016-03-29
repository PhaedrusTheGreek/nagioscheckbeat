# NagiosCheckBeat

NagiosCheckBeat is the [Beat](https://www.elastic.co/products/beats) used for
running Nagios checks.    You can integrate with [Watcher](https://www.elastic.co/blog/a-case-for-self-monitoring-systems) or [Nagios Core](https://discuss.elastic.co/t/questions-about-self-monitoring-systems-blog-post/43542/12?u=phaedrusthegreek) for Alerting.

![Kibana Screenshot](https://github.com/PhaedrusTheGreek/nagioscheckbeat/blob/master/ss.png)

## Template

To apply the ES Template:

```
curl -XPUT 'http://localhost:9200/_template/nagioscheckbeat' -d@etc/nagioscheckbeat.template.json
```

## Configuration
```
############################# Input ############################################
input:
  checks:
    -
      name: "heartbeat"
      cmd: "/usr/lib64/nagios/plugins/check_dummy"
      args: "0 Checking In!"
      period: "10s"
    -
      name: "disks"
      cmd: "/usr/lib64/nagios/plugins/check_disk"
      args: "-w 80 -c 90 -x /dev"
      period: "1h"
    -
      name: "load"
      cmd: "/usr/lib64/nagios/plugins/check_load"
      args: "-w 5 -c 10"
      period: "1m"
    -
      name: "io"
      cmd: "/usr/lib64/nagios/plugins/check_sar_perf.py"
      args: "io_transfer"
      period: "30s"
      enabled: false
```

## Produces

Firstly, the metrics, individually as documents:

```
{
   "took": 1,
   "timed_out": false,
   "_shards": {
      "total": 1,
      "successful": 1,
      "failed": 0
   },
   "hits": {
      "total": 12,
      "max_score": 1,
      "hits": [
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bdpWMO5Jb2lqbx",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.924Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 50178,
               "label": "/",
               "max": 50268,
               "min": 0,
               "name": "disks",
               "type": "nagiosmetric",
               "uom": "MB",
               "value": 33124,
               "warning": 50188
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bdpWMO5Jb2lqby",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.924Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 3910,
               "label": "/dev/shm",
               "max": 4000,
               "min": 0,
               "name": "disks",
               "type": "nagiosmetric",
               "uom": "MB",
               "value": 0,
               "warning": 3920
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bdpWMO5Jb2lqbz",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.924Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 386,
               "label": "/boot",
               "max": 476,
               "min": 0,
               "name": "disks",
               "type": "nagiosmetric",
               "uom": "MB",
               "value": 154,
               "warning": 396
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bdpWMO5Jb2lqb0",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.924Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 76800,
               "label": "/home",
               "max": 76890,
               "min": 0,
               "name": "disks",
               "type": "nagiosmetric",
               "uom": "MB",
               "value": 363,
               "warning": 76810
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bdpWMO5Jb2lqb2",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.933Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 10,
               "label": "load1",
               "max": 0,
               "min": 0,
               "name": "load",
               "type": "nagiosmetric",
               "uom": "",
               "value": 0.16,
               "warning": 5
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bepWMO5Jb2lqb3",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.933Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 10,
               "label": "load5",
               "max": 0,
               "min": 0,
               "name": "load",
               "type": "nagiosmetric",
               "uom": "",
               "value": 0.05,
               "warning": 5
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7bepWMO5Jb2lqb4",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.933Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 10,
               "label": "load15",
               "max": 0,
               "min": 0,
               "name": "load",
               "type": "nagiosmetric",
               "uom": "",
               "value": 0.06,
               "warning": 5
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7pTpWMO5Jb2lqb6",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.948Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 0,
               "label": "tps",
               "max": 0,
               "min": 0,
               "name": "io",
               "type": "nagiosmetric",
               "uom": "",
               "value": 0,
               "warning": 0
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7pTpWMO5Jb2lqb7",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.948Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 0,
               "label": "rtps",
               "max": 0,
               "min": 0,
               "name": "io",
               "type": "nagiosmetric",
               "uom": "",
               "value": 0,
               "warning": 0
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVH0P7pTpWMO5Jb2lqb8",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.948Z",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "count": 1,
               "critical": 0,
               "label": "wtps",
               "max": 0,
               "min": 0,
               "name": "io",
               "type": "nagiosmetric",
               "uom": "",
               "value": 0,
               "warning": 0
            }
         }
      ]
   }
}
```

Secondly, the results of the actual Nagios Checks, as a separate *type*

```
{
   "took": 1,
   "timed_out": false,
   "_shards": {
      "total": 1,
      "successful": 1,
      "failed": 0
   },
   "hits": {
      "total": 4,
      "max_score": 1,
      "hits": [
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagioscheck",
            "_id": "AVH0P7bdpWMO5Jb2lqbv",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.922Z",
               "args": "0 Hello",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_dummy",
               "count": 1,
               "message": "OK: Hello\n",
               "status": "OK",
               "took_ms": 2,
               "type": "nagioscheck"
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagioscheck",
            "_id": "AVH0P7bdpWMO5Jb2lqbw",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.924Z",
               "args": "-w 80 -c 90 -x /dev",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_disk",
               "count": 1,
               "message": "DISK OK - free space: / 14568 MB (30% inode=94%); /dev/shm 4000 MB (100% inode=99%); /boot 292 MB (65% inode=99%); /home 72598 MB (99% inode=99%);",
               "status": "OK",
               "took_ms": 8,
               "type": "nagioscheck"
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagioscheck",
            "_id": "AVH0P7bdpWMO5Jb2lqb1",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.933Z",
               "args": "-w 5 -c 10",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_load",
               "count": 1,
               "message": "OK - load average: 0.16, 0.05, 0.06",
               "status": "OK",
               "took_ms": 14,
               "type": "nagioscheck"
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagioscheck",
            "_id": "AVH0P7pTpWMO5Jb2lqb5",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T18:56:33.948Z",
               "args": "io_transfer",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_sar_perf.py",
               "count": 1,
               "message": "sar OK ",
               "status": "OK",
               "took_ms": 1062,
               "type": "nagioscheck"
            }
         }
      ]
   }
}
```



