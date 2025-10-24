#!/bin/bash

cd "$(dirname "$0")/.."
first=4
cnt=100
# rm -rf ./metrics_data/fiber_spike_deep/* ./metrics_data/fiber_spike_flat/* ./metrics_data/gin_spike_deep/* ./metrics_data/gin_spike_flat/*


mkdir -p ./metrics_data/fiber_spike_deep ./metrics_data/fiber_spike_flat ./metrics_data/gin_spike_deep ./metrics_data/gin_spike_flat
docker compose -f ./deployment/docker-compose.yml stop fiber-app
docker compose -f ./deployment/docker-compose.yml stop gin-app

start_script=$(date)

for ((i=first; i<=cnt; i++))
do
    echo "i = $i"
    start_iteration=$(date)
    echo "Start iteration: $start_iteration" 

    savefile="fiber_spike_deep"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml start fiber-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
    pandora ./requests/pandora_config/deep_spike.yaml > /dev/null 2>&1
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    make dump_logs
    mkdir -p "./metrics_data/${savefile}/${i}"                                          # mkdir -p "./metrics_data/fiber_spike_deep/1"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"       # mv ./metrics_data/logs/logs_time_series.txt ./metrics_data/fiber_spike_deep/1
    # end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end" -container="deployment-fiber-app-1"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
    docker compose -f ./deployment/docker-compose.yml stop fiber-app

    savefile="fiber_spike_flat"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml start fiber-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
    pandora ./requests/pandora_config/flat_spike.yaml > /dev/null 2>&1
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    make dump_logs
    mkdir -p "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    # end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end" -container="deployment-fiber-app-1"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
    docker compose -f ./deployment/docker-compose.yml stop fiber-app

    savefile="gin_spike_deep"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)   
    pandora ./requests/pandora_config/deep_spike.yaml > /dev/null 2>&1
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    make dump_logs
    mkdir -p "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    # end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end" -container="deployment-gin-app-1"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
    docker compose -f ./deployment/docker-compose.yml stop gin-app

    savefile="gin_spike_flat"
    echo "$savefile"
    docker compose -f ./deployment/docker-compose.yml start gin-app
    sleep 1
    start=$(date -u +%Y-%m-%dT%H:%M:%SZ)    # В RFC3339 формате
    pandora ./requests/pandora_config/flat_spike.yaml > /dev/null 2>&1
    end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    make dump_logs
    mkdir -p "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/logs/logs_time_series.txt "./metrics_data/${savefile}/${i}"
    mv ./metrics_data/pandora_results/pandora.phout "./metrics_data/${savefile}/${i}"
    # end=$(date -u +%Y-%m-%dT%H:%M:%SZ)
    go run ./cmd/export_prom/main.go -start="$start" -end="$end" -container="deployment-gin-app-1"
    mv ./metrics_data/prometheus/* "./metrics_data/${savefile}/${i}"
    docker compose -f ./deployment/docker-compose.yml stop gin-app

    echo "Start iteration: $start_iteration" 
    end_iteration=$(date)
    echo "End iteration: $end_iteration" 
    start_seconds=$(date -d "$start_iteration" +%s)
    end_seconds=$(date -d "$end_iteration" +%s)
    duration=$((end_seconds - start_seconds))
    echo "Duration_iteration: $duration" 
done

end_script=$(date)

echo "Start script: $start_script" 
echo "End script: $end_script" 

start_seconds=$(date -d "$start_script" +%s)
end_seconds=$(date -d "$end_script" +%s)
duration=$((end_seconds - start_seconds))
echo "Duration_script: $duration" 
