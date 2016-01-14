FROM centos:latest

RUN yum -y update && yum clean all
RUN yum -y install curl nc httpd epel-release
RUN yum -y install nagios-plugins-all perl-LWP-Protocol-https.noarch perl-Nagios-Plugin.noarch
RUN yum clean all
RUN rpm -ivh https://github.com/PhaedrusTheGreek/nagioscheckbeat/blob/master/build/nagioscheckbeat-0.5.3-x86_64.rpm?raw=true

EXPOSE 80

COPY httpd.conf /etc/httpd/conf/httpd.conf
COPY nagioscheckbeat.yml /etc/nagioscheckbeat/nagioscheckbeat.yml

COPY check_apachestatus.pl /usr/lib64/nagios/plugins/check_apachestatus.pl
RUN chmod +x /usr/lib64/nagios/plugins/check_apachestatus.pl

ADD run-stuff.sh /run-stuff.sh
RUN chmod -v +x /run-stuff.sh

CMD ["/run-stuff.sh"]
