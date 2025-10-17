#!/bin/bash

cd "$(dirname "$0")/.."
# make down_all
# make run_prom

# make down_gin_app
# make run_gin_app 
# sleep 2
# make pandora
# make dump_logs

for i in {1..10}
do
    docker compose -f ./deployment/docker-compose.yml stop gin-app
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    make pandora
    make dump_logs
done