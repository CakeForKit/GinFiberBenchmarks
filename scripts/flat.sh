#!/bin/bash

cd "$(dirname "$0")/.."
# make down_all
# make run_prom

make down_gin_app
make run_gin_app 
sleep 2
make pandora
make dump_logs

