#!/bin/bash

echo "Testing POST /helloworld endpoint..."
echo ""

curl -X POST http://localhost:6000/helloworld \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from GinHub!"}' \
  | jq .

echo ""
echo "Test completed!"
