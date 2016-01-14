#!/bin/bash
echo "Putting ES Template..."
curl -s -XPUT 'http://127.0.0.1:9200/_template/nagioscheckbeat' -d@../etc/nagioscheckbeat.template.json
