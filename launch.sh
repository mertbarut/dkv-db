#!/bin/bash
set -e

cd $(dirname $0)

go run main.go -db-location=amsterdam.db -http-addr=127.0.0.1:8080 -config-file=sharding.toml -shard=Amsterdam &
go run main.go -db-location=rotterdam.db -http-addr=127.0.0.1:8081 -config-file=sharding.toml -shard=Rotterdam &
go run main.go -db-location=groningen.db -http-addr=127.0.0.1:8082 -config-file=sharding.toml -shard=Groningen &

# Testing redirection between shards

sleep 1

curl 'http://localhost:8080/set?key=42Wolfsburg&value=Good' &
curl 'http://localhost:8080/get?key=42Wolfsburg' &

wait