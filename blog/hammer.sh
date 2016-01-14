#!/bin/bash

# Requires mysql
docker exec blog_mysql_1 mysql -e"create database mysqlslap"
mysqlslap -v -h127.0.0.1  –auto-generate-sql --iterations=100 -concurrency=200 --auto-generate-sql --verbose >> /tmp/hammer.log 2>&1 &

# Requires apache
ab -n 100000 -c 200 http://localhost/ >> /tmp/hammer.log 2>&1 &

# Requires redis
redis-benchmark -c 200 -n 1000000 >> /tmp/hammer.log 2>&1 &

read -p "Hammering... [Enter] key to stop...  (See log in /tmp/hammer.log)"

killall redis-benchmark
killall ab
killall mysqlslap
