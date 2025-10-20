#!/bin/bash

cd "$(dirname "$0")/.."
first=5
cnt=100

# rm -rf ./metrics_data/gin_deep_ramp_up/* ./metrics_data/gin_flat_ramp_up/* ./metrics_data/fiber_deep_ramp_up/* ./metrics_data/fiber_flat_ramp_up/*


for ((i=first; i<=cnt; i++))
do
    echo "i = $i"

    savefile="fiber_deep_ramp_up"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml stop fiber-app
    docker compose -f ./deployment/docker-compose.yml start fiber-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
    pandora ./requests/pandora_config/deep_ramp_up.yaml > /dev/null 2>&1
    make dump_logs
    mkdir "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"


    savefile="fiber_flat_ramp_up"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml stop fiber-app
    docker compose -f ./deployment/docker-compose.yml start fiber-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
    pandora ./requests/pandora_config/flat_ramp_up.yaml > /dev/null 2>&1
    make dump_logs
    mkdir "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"


    savefile="gin_deep_ramp_up"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml stop gin-app
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
    pandora ./requests/pandora_config/deep_ramp_up.yaml > /dev/null 2>&1
    make dump_logs
    mkdir "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"


    savefile="gin_flat_ramp_up"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml stop gin-app
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
    pandora ./requests/pandora_config/flat_ramp_up.yaml > /dev/null 2>&1
    make dump_logs
    mkdir "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
done



# savefile="fiber_deep_ramp_up"
# for ((i=first; i<=cnt; i++))
# do
#     docker compose -f ./deployment/docker-compose.yml stop fiber-app
#     docker compose -f ./deployment/docker-compose.yml start fiber-app
#     sleep 1
#     start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
#     pandora ./requests/pandora_config/deep_ramp_up.yaml
    
#     make dump_logs
#     mkdir "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
#     end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
#     go run ./cmd/export_prom/main.go -start="$start" -end="$end"
#     mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
# done


# savefile="fiber_flat_ramp_up"
# # rm -rf ./metrics_data/fiber_flat_ramp_up/*
# for ((i=first; i<=cnt; i++))
# do
#     docker compose -f ./deployment/docker-compose.yml stop fiber-app
#     docker compose -f ./deployment/docker-compose.yml start fiber-app
#     sleep 1
#     start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
#     pandora ./requests/pandora_config/flat_ramp_up.yaml
    
#     make dump_logs
#     mkdir "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
#     end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
#     go run ./cmd/export_prom/main.go -start="$start" -end="$end"
#     mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
# done


# savefile="gin_deep_ramp_up"
# # rm -rf ./metrics_data/gin_deep_ramp_up/*
# for ((i=first; i<=cnt; i++))
# do
#     docker compose -f ./deployment/docker-compose.yml stop gin-app
#     docker compose -f ./deployment/docker-compose.yml start gin-app
#     sleep 1
#     start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
#     pandora ./requests/pandora_config/deep_ramp_up.yaml
    
#     make dump_logs
#     mkdir "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
#     end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
#     go run ./cmd/export_prom/main.go -start="$start" -end="$end"
#     mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
# done


# savefile="gin_flat_ramp_up"
# # rm -rf ./metrics_data/gin_flat_ramp_up/*
# for ((i=first; i<=cnt; i++))
# do
#     docker compose -f ./deployment/docker-compose.yml stop gin-app
#     docker compose -f ./deployment/docker-compose.yml start gin-app
#     sleep 1
#     start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
#     pandora ./requests/pandora_config/flat_ramp_up.yaml
    
#     make dump_logs
#     mkdir "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
#     mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
#     end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
#     go run ./cmd/export_prom/main.go -start="$start" -end="$end"
#     mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
# done
