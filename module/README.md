Integration with Metric Beats requires that the check library be imported

### Import the Check (in metricbeat beat/list.go):

```
import (
        _ "github.com/PhaedrusTheGreek/nagioscheckbeat/module/check"
)
```

### Configure the Module and the Check

```
metricbeat:
  modules:
    nagioscheck:
      metrics:
        check:
            period: 10s
            name: "heartbeat"
            cmd: "/usr/local/sbin/check_dummy"
            args: "0 Hello"
```
