#!/bin/sh

SERVER="127.0.0.1:9999"

TOKEN=$(curl --silent -X POST http://${SERVER}/api/login --data '{"username": "test", "password": "passwd"}' | jq -rc .token)

curl -v -H "Authorization: Bearer ${TOKEN}" http://127.0.0.1/private/index.html
