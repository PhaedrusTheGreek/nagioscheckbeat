#!/bin/bash

cd /tmp/
wget --no-check-certificate https://labs.consol.de/assets/downloads/nagios/check_mysql_health-2.2.1.tar.gz
tar xvzf check_mysql_health-2.2.1.tar.gz 
cd /tmp/check_mysql_health-2.2.1
./configure --prefix=/usr/lib/nagios
make install

