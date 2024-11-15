#!/bin/bash

# Install wrk if not installed
if ! command -v wrk &> /dev/null
then
    echo "wrk could not be found, please install it first."
    exit
fi

# Benchmark parameters
CONCURRENT_CONNECTIONS=1000
DURATION=60s
THREADS=10

# Start benchmark
wrk -t$THREADS -c$CONCURRENT_CONNECTIONS -d$DURATION -s scripts/benchmark.lua http://localhost:8080/stream/start