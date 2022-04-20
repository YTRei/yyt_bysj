#! /bin/bash

#Build web and other services
cd D:/study/bysj/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd D:/study/bysj/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd D:/study/bysj/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd D:/study/bysj/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web
