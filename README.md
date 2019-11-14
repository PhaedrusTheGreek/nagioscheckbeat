# NagiosCheckBeat

NagiosCheckBeat is the [Beat](https://www.elastic.co/products/beats) used for
running Nagios checks.    

Check out [this blog post](https://www.elastic.co/blog/a-case-for-self-monitoring-systems) on how it works.

You can integrate with Watcher [(How-To)](https://www.elastic.co/blog/a-case-for-self-monitoring-systems) or Nagios Core [(How-To)](https://discuss.elastic.co/t/questions-about-self-monitoring-systems-blog-post/43542/12?u=phaedrusthegreek) for Alerting.

![Kibana Screenshot](https://github.com/PhaedrusTheGreek/nagioscheckbeat/blob/master/ss.png)

## Compatibility

- For Elasticsearch 1.x compatibility, see the 0.5.x branch.
- NagiosCheckBeat 0.6.0 is compatible with Elasticsaerch 2.x through 5.x.   
- NagiosCheckbeat 6.2.3 is compatible with Elastic Stack 6.x 
- NagiosCheckBeat 7.6.0 is compatible with Elastic Stack 7.x

## Security

- As of NagiosCheckBeat 7.x, process forking must be explicitly allowed by disabling seccomp.  Do note the security hazard here!  Keep your config secure!  

```
seccomp.enabled: false
```

## Download & Install

Packages for your OS can be found in [Releases](https://github.com/PhaedrusTheGreek/nagioscheckbeat/releases)

For example, in an i686 architecture:

```
$ sudo rpm -i https://github.com/PhaedrusTheGreek/nagioscheckbeat/releases/download/6.2.3/nagioscheckbeat-6.2.3-i686.rpm
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

Firstly, the performance data metrics, individually as documents `type: nagiosmetric`:

```
{
  "took": 3,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": 112,
    "max_score": 1.6964493,
    "hits": [
      {
        "_index": "nagioscheckbeat-6.2.2-2018.02.20",
        "_type": "doc",
        "_id": "-39LtGEBVAhdWhA-GWkB",
        "_score": 1.6964493,
        "_source": {
          "@timestamp": "2018-02-20T17:37:54.953Z",
          "value": 638,
          "beat": {
            "name": "Jasons-MacBook-Pro-912.local",
            "hostname": "Jasons-MacBook-Pro-912.local",
            "version": "6.2.2"
          },
          "uom": "",
          "max": 0,
          "critical": 0,
          "min": 0,
          "name": "proc",
          "label": "procs",
          "warning": 0,
          "type": "nagiosmetric"
        },
        "fields": {
          "name": [
            "proc"
          ]
        }
      },
      {
        "_index": "nagioscheckbeat-6.2.2-2018.02.20",
        "_type": "doc",
        "_id": "O39KtGEBVAhdWhA-o2nY",
        "_score": 1.6964493,
        "_source": {
          "@timestamp": "2018-02-20T17:37:24.949Z",
          "name": "load",
          "uom": "",
          "warning": 5,
          "label": "load5",
          "critical": 10,
          "max": 0,
          "type": "nagiosmetric",
          "beat": {
            "name": "Jasons-MacBook-Pro-912.local",
            "hostname": "Jasons-MacBook-Pro-912.local",
            "version": "6.2.2"
          },
          "value": 0,
          "min": 0
        },
        "fields": {
          "name": [
            "load"
          ]
        }
      }
    ]
  }
}
```

Secondly, the results of the actual Nagios Checks, where `type: nagioscheck`

```
{
  "took": 2,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": 341,
    "max_score": 0.38229805,
    "hits": [
      {
        "_index": "nagioscheckbeat-6.2.2-2018.02.20",
        "_type": "doc",
        "_id": "s39KtGEBVAhdWhA-6mkg",
        "_score": 0.38229805,
        "_source": {
          "@timestamp": "2018-02-20T17:37:42.953Z",
          "took_ms": 11,
          "type": "nagioscheck",
          "name": "heartbeat",
          "message": "OK: Hello\n",
          "beat": {
            "name": "Jasons-MacBook-Pro-912.local",
            "hostname": "Jasons-MacBook-Pro-912.local",
            "version": "6.2.2"
          },
          "status": "OK",
          "status_code": 0,
          "cmd": "/usr/local/sbin/check_dummy",
          "args": "0 Hello"
        },
        "fields": {
          "name": [
            "heartbeat"
          ]
        }
      },
      {
        "_index": "nagioscheckbeat-6.2.2-2018.02.20",
        "_type": "doc",
        "_id": "u39KtGEBVAhdWhA-8Wnv",
        "_score": 0.38229805,
        "_source": {
          "@timestamp": "2018-02-20T17:37:44.952Z",
          "cmd": "/usr/local/sbin/check_procs",
          "args": "",
          "message": "PROCS OK: 638 processes ",
          "status_code": 0,
          "took_ms": 86,
          "beat": {
            "name": "Jasons-MacBook-Pro-912.local",
            "hostname": "Jasons-MacBook-Pro-912.local",
            "version": "6.2.2"
          },
          "status": "OK",
          "type": "nagioscheck",
          "name": "proc"
        },
        "fields": {
          "name": [
            "proc"
          ]
        }
      },
      {
        "_index": "nagioscheckbeat-6.2.2-2018.02.20",
        "_type": "doc",
        "_id": "9X9LtGEBVAhdWhA-GWkB",
        "_score": 0.38229805,
        "_source": {
          "@timestamp": "2018-02-20T17:37:54.949Z",
          "cmd": "/usr/local/sbin/check_load",
          "status_code": 0,
          "took_ms": 14,
          "type": "nagioscheck",
          "name": "load",
          "args": "-w 5 -c 10",
          "message": "OK - load average: 2.19, 0.00, 0.00",
          "status": "OK",
          "beat": {
            "hostname": "Jasons-MacBook-Pro-912.local",
            "version": "6.2.2",
            "name": "Jasons-MacBook-Pro-912.local"
          }
        },
        "fields": {
          "name": [
            "load"
          ]
        }
      },
      {
        "_index": "nagioscheckbeat-6.2.2-2018.02.20",
        "_type": "doc",
        "_id": "6X9OtGEBVAhdWhA-Jm5P",
        "_score": 0.30196804,
        "_source": {
          "@timestamp": "2018-02-20T17:41:14.956Z",
          "name": "disks",
          "cmd": "/usr/local/sbin/check_disk",
          "message": "DISK CRITICAL - /Users/jason/OBFUSCATED is not accessible: No such file or directory\n",
          "beat": {
            "name": "Jasons-MacBook-Pro-912.local",
            "hostname": "Jasons-MacBook-Pro-912.local",
            "version": "6.2.2"
          },
          "status": "CRITICAL",
          "status_code": 2,
          "took_ms": 27,
          "type": "nagioscheck",
          "args": "-w 80 -c 90"
        },
        "fields": {
          "name": [
            "disks"
          ]
        }
      }
    ]
  }
}
```



