#!/bin/bash

/usr/bin/nagioscheckbeat -c /etc/nagioscheckbeat/nagioscheckbeat.yml &

/usr/bin/redis-server
