# NagiosCheckBeat

NagiosCheckBeat is the [Beat](https://www.elastic.co/products/beats) used for
Running Nagios Checks.


## Template

To apply the ES Template:

```
curl -XPUT 'http://localhost:9200/_template/nagioscheckbeat' -d@etc/nagioscheckbeat.template.json
```

## Configuration
```
############################# Input ############################################
input:
  interval: "5s"
  checks:
    -
      name: "disks"
      cmd: "/usr/local/sbin/check_disk"
      args: "-w 80 -c 90 -x /dev"
    -
      name: "load"
      cmd: "/usr/local/sbin/check_load"
      args: "-w 5 -c 10"
```

## Produces
```
{
  "@timestamp": "2015-12-29T18:24:54.717Z",
  "args": "-w 80 -c 90 -x /dev",
  "beat": {
    "hostname": "ptg-mbp",
    "name": "ptg-mbp"
  },
  "cmd": "/usr/local/sbin/check_disk",
  "count": 1,
  "disks": {
    "/": {
      "Label": "/",
      "Uom": "MB",
      "Value": 88051,
      "Warning": 114464,
      "Critical": 114454,
      "Min": 0,
      "Max": 114544
    }
  },
  "message": "DISK OK - free space: / 26242 MB (22% inode=22%);",
  "status": "OK",
  "took_ms": 12,
  "type": "nagioscheck"
}

{
  "@timestamp": "2015-12-29T18:24:54.729Z",
  "args": "-w 5 -c 10",
  "beat": {
    "hostname": "ptg-mbp",
    "name": "ptg-mbp"
  },
  "cmd": "/usr/local/sbin/check_load",
  "count": 1,
  "load": {
    "load1": {
      "Label": "load1",
      "Uom": "",
      "Value": 2.35,
      "Warning": 5,
      "Critical": 10,
      "Min": 0,
      "Max": 0
    },
    "load15": {
      "Label": "load15",
      "Uom": "",
      "Value": 0,
      "Warning": 5,
      "Critical": 10,
      "Min": 0,
      "Max": 0
    },
    "load5": {
      "Label": "load5",
      "Uom": "",
      "Value": 0,
      "Warning": 5,
      "Critical": 10,
      "Min": 0,
      "Max": 0
    }
  },
  "message": "OK - load average: 2.35, 0.00, 0.00",
  "status": "OK",
  "took_ms": 12,
  "type": "nagioscheck"
}
```
