#!/usr/bin/env bash

URL="http://localhost:4000"
COUNT=300  # number of requests to send

echo "Sending $COUNT requests to $URL..."

for i in $(seq 1 $COUNT); do
    RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" "$URL")
    echo "Request $i → HTTP $RESPONSE"
done