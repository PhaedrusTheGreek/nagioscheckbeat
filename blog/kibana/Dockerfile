FROM kibana:latest

RUN apt-get update && apt-get install -y netcat

COPY run-stuff.sh /run-stuff.sh
RUN chmod +x /run-stuff.sh

RUN kibana plugin --install elastic/sense
RUN kibana plugin --install elastic/timelion

CMD ["/run-stuff.sh"]
