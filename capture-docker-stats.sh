#!/bin/bash

# Set CSV file to log stats to  
STATS_FILE=docker-stats.csv

# Interval in seconds to capture stats
INTERVAL=1

# Check if trigger file exists
while [ ! -f /tmp/triggerscript ]; do
  sleep 1
done

# Header for CSV  
echo -e "Name,CPU%,Memory Usage" > $STATS_FILE

# Capture stats continuously 
while true; do  
  
  STATS=$(docker stats --no-stream --format \
           "{{.Name}},{{.CPUPerc}},{{.MemUsage}}")
    
  # Append to CSV
  echo "$STATS" >> $STATS_FILE

  sleep $INTERVAL
  
done