### A case for self monitoring systems

https://www.elastic.co/blog/a-case-for-self-monitoring-systems

This is a proof of concept that uses nagioscheckbeat in combination with Elasticsearch, Kibana, and Watcher to implement a total monitoring solution.  Tested on CentOS 6.

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

