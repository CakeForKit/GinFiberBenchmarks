#!/bin/bash

cd "$(dirname "$0")/.."
cnt=2
savefile="gin_deep_ramp_up"
rm -rf ./metrics_data/gin_deep_ramp_up/*
for ((i=1; i<=cnt; i++))
do
    docker compose -f ./deployment/docker-compose.yml stop gin-app
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
    pandora ./requests/pandora_config/deep_ramp_up.yaml
    
    make dump_logs
    mkdir "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
done


savefile="gin_flat_ramp_up"
# rm -rf ./metrics_data/gin_flat_ramp_up/*
for ((i=1; i<=cnt; i++))
do
    docker compose -f ./deployment/docker-compose.yml stop gin-app
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
    pandora ./requests/pandora_config/flat_ramp_up.yaml
    
    make dump_logs
    mkdir "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
done
