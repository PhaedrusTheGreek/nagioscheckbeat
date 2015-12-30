# NagiosCheckBeat

NagiosCheckBeat is the [Beat](https://www.elastic.co/products/beats) used for
Running Nagios Checks.

1 document is published for each check, and then another 1 document for each metric in the check results.

## Template

To apply the ES Template:

```
curl -XPUT 'http://localhost:9200/_template/nagioscheckbeat' -d@etc/nagioscheckbeat.template.json
```

## Configuration
```
############################# Input ############################################
input:
  interval: "30s"
  checks:
    -
      name: "heartbeat"
      cmd: "/usr/local/sbin/check_dummy"
      args: "0 Checking In!"
    -
      name: "disks"
      cmd: "/usr/local/sbin/check_disk"
      args: "-w 80 -c 90 -x /dev"
    -
      name: "load"
      cmd: "/usr/local/sbin/check_load"
      args: "-w 5 -c 10"
    -
      name: "io"
      cmd: "/usr/lib64/nagios/plugins/check_sar_perf.py"
      args: "io_transfer"
```

## Produces

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
      "total": 15,
      "max_score": 1,
      "hits": [
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagioscheck",
            "_id": "AVHz_44dpWMO5Jb2lqbP",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.194Z",
               "args": "-w 80 -c 90 -x /dev",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_disk",
               "count": 1,
               "message": "DISK OK - free space: / 14568 MB (30% inode=94%); /dev/shm 4000 MB (100% inode=99%); /boot 292 MB (65% inode=99%); /home 72598 MB (99% inode=99%);",
               "status": "OK",
               "took_ms": 2,
               "type": "nagioscheck"
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVHz_44dpWMO5Jb2lqbQ",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.194Z",
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
            "_id": "AVHz_44dpWMO5Jb2lqbR",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.194Z",
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
            "_id": "AVHz_44dpWMO5Jb2lqbS",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.194Z",
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
            "_id": "AVHz_44dpWMO5Jb2lqbT",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.194Z",
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
            "_type": "nagioscheck",
            "_id": "AVHz_44dpWMO5Jb2lqbU",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.196Z",
               "args": "-w 5 -c 10",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_load",
               "count": 1,
               "message": "OK - load average: 0.00, 0.01, 0.05",
               "status": "OK",
               "took_ms": 4,
               "type": "nagioscheck"
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVHz_44dpWMO5Jb2lqbV",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.196Z",
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
               "value": 0,
               "warning": 5
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVHz_44dpWMO5Jb2lqbW",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.196Z",
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
               "value": 0.01,
               "warning": 5
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagiosmetric",
            "_id": "AVHz_44dpWMO5Jb2lqbX",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.196Z",
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
               "value": 0.05,
               "warning": 5
            }
         },
         {
            "_index": "nagioscheckbeat-2015.12.30",
            "_type": "nagioscheck",
            "_id": "AVHz_5GZpWMO5Jb2lqbY",
            "_score": 1,
            "_source": {
               "@timestamp": "2015-12-30T17:46:29.200Z",
               "args": "io_transfer",
               "beat": {
                  "hostname": "max.elastic.co",
                  "name": "max.elastic.co"
               },
               "cmd": "/usr/lib64/nagios/plugins/check_sar_perf.py",
               "count": 1,
               "message": "sar OK ",
               "status": "OK",
               "took_ms": 1039,
               "type": "nagioscheck"
            }
         }
      ]
   }
}
```
