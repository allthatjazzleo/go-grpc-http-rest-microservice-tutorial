#! /bin/sh
./server -grpc-port=9090 -db-host=localhost:3306 -http-port=8080 -db-user=leo -db-password=12345678 -db-schema=Grpc -log-level=-1 -log-time-format=2006-01-02T15:04:05.999999999Z07:00