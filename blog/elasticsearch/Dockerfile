FROM elasticsearch:latest

ENV ES_HEAP_SIZE 4G

EXPOSE 9200

RUN /usr/share/elasticsearch/bin/plugin install license
RUN /usr/share/elasticsearch/bin/plugin install watcher

ENTRYPOINT ["/docker-entrypoint.sh"]
