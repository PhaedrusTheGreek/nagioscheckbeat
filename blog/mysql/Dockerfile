FROM orchardup/mysql

RUN apt-get update 
RUN apt-get install -y nagios-plugins
RUN apt-get install -y vim wget make

ENV URL 'https://github.com/PhaedrusTheGreek/nagioscheckbeat/blob/master/build/nagioscheckbeat_0.5.3_amd64.deb?raw=true'
ENV FILE /tmp/tmp-file.deb	
RUN wget --no-check-certificate "$URL" -qO $FILE && dpkg -i $FILE; rm $FILE

EXPOSE 3306

COPY nagioscheckbeat.yml /etc/nagioscheckbeat/nagioscheckbeat.yml

ADD install-check.sh /install-check.sh
RUN chmod -v +x /install-check.sh
RUN /install-check.sh

ADD run-stuff.sh /run-stuff.sh
RUN chmod -v +x /run-stuff.sh

CMD ["/run-stuff.sh"]
