### A case for self monitoring systems

This is a proof of concept that uses nagioscheckbeat in combination with Elasticsearch, Kibana, and Watcher to implement a total monitoring solution.  Tested on CentOS 6.

Requires:
- Docker

Optional Requirements
- Load Test (hammer.sh) requires redis, mysql, and apache.

### Instructions

1. Set your gmail username and password in elasticearch/config/elasticsearch.yml
2. Run Stuff

```
docker-compose build
docker-compose up -d elasticsearch
./put-template.sh
docker-compose up 
```
