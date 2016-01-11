FROM centos:latest
MAINTAINER http://www.centos.org
LABEL Vendor="CentOS"
LABEL License=GPLv2
LABEL Version=2.4.6-31

RUN yum -y update && yum clean all
RUN yum -y install httpd epel-release
RUN yum -y install nagios-plugins-all perl-LWP-Protocol-https.noarch perl-Nagios-Plugin.noarch
RUN yum clean all
RUN rpm -ivh https://github.com/PhaedrusTheGreek/nagioscheckbeat/blob/master/build/nagioscheckbeat-0.5.0-x86_64.rpm?raw=true

EXPOSE 80

# Simple startup script to avoid some issues observed with container restart 
ADD run-httpd.sh /run-httpd.sh
RUN chmod -v +x /run-httpd.sh

# Config, which enables server-status
COPY httpd.conf /etc/httpd/conf/httpd.conf
COPY nagioscheckbeat.yml /etc/nagioscheckbeat/nagioscheckbeat.yml

# Nagios Check
COPY check_apachestatus.pl /usr/lib64/nagios/plugins/check_apachestatus.pl
RUN chmod +x /usr/lib64/nagios/plugins/check_apachestatus.pl

CMD ["/run-httpd.sh"]
