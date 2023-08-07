#!/bin/bash

WD=$(cd $(dirname $0); pwd -P)

docker rm -f testdb-mysql

docker run -d \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=testdb \
  -e MYSQL_USER=testdb \
  -e MYSQL_PASSWORD=testdb \
  -p 3306:3306 \
  -v ${WD}/testdata:/docker-entrypoint-initdb.d \
  --name testdb-mysql mysql:8.0
