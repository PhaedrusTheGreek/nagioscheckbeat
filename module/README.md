metricbeat:
  modules:
    nagioscheck:
      metrics:
        check:
            period: 10s
            name: "heartbeat"
            cmd: "/usr/local/sbin/check_dummy"
            args: "0 Hello"
