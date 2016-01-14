FROM ubuntu

RUN apt-get update && apt-get install -y redis-server wget 
RUN apt-get install -y nagios-plugins 
RUN apt-get install -y libredis-perl

ENV URL 'https://github.com/PhaedrusTheGreek/nagioscheckbeat/blob/master/build/nagioscheckbeat_0.5.3_amd64.deb?raw=true'
ENV FILE /tmp/tmp-file.deb	
RUN wget "$URL" -qO $FILE && dpkg -i $FILE; rm $FILE

EXPOSE 6379

COPY nagioscheckbeat.yml /etc/nagioscheckbeat/nagioscheckbeat.yml

COPY check_redis.pl /usr/lib/nagios/plugins/check_redis.pl
RUN chmod +x /usr/lib/nagios/plugins/check_redis.pl

ADD run-stuff.sh /run-stuff.sh
RUN chmod -v +x /run-stuff.sh

CMD ["/run-stuff.sh"]
