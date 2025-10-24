#!/bin/bash
start_script=$(date)
echo "start_script: $start_script" 
echo "End script: $(date)" 
sleep 1
end=$(date)
start_seconds=$(date -d "$start_script" +%s)
end_seconds=$(date -d "$end" +%s)
duration=$((end_seconds - start_seconds))
echo "Duration: $duration" 