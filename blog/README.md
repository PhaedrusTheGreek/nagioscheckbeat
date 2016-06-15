### A case for self monitoring systems

https://www.elastic.co/blog/a-case-for-self-monitoring-systems

This is a proof of concept that uses nagioscheckbeat in combination with Elasticsearch, Kibana, and Watcher to implement a total monitoring solution.  Tested on CentOS 6.

##### Errata
- In *Alerting On Lost Heartbeats* , a *Threshold* parameter is sent to the script, but never used
- A better solution was found later found by [inqueue](https://github.com/inqueue) - Re: Alerting On Lost Heartbeats, instead of using a groovy script for the watch to determine host *downness*, you could do [something like this](https://gist.github.com/inqueue/24b459c177bc0e1967198008eb3a40d4), where we aggregate on heartbeats in the last now-30s seconds.  Since the parent query searches for now-1d, we end up with an empty bucket for down hosts.   

### Requirements

Requires:
- Docker 

Optional Requirements
- Load Test (hammer.sh) requires redis, mysql, and apache.

### Instructions

1. Set your gmail username and password in elasticearch/config/elasticsearch.yml (Also see [here](https://support.google.com/accounts/answer/6010255?hl=en), and make sure 2 phase auth is not enabled)
2. Run Stuff

  ```
  docker-compose build
  docker-compose up -d elasticsearch
  ./put-template.sh
  docker-compose up 
  ```

3. Navigate to Kibana @ http://docker-server:5601 and configure an index pattern of `nagioscheckbeat*`.  Then load the `kibana-dashboard.json` into Objects -> Import.

4. run `hammer.sh` to generate something interesting

