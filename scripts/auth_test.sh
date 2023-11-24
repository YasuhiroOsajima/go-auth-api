#!/bin/sh

SERVER="127.0.0.1:9999"

TOKEN=$(curl --silent -X POST http://${SERVER}/api/login --data '{"username": "test", "password": "passwd"}' | jq -rc .token)

echo "NG patter test:"
curl -v http://${SERVER}/api/admin/user
echo ""

echo ""
echo "OK pattern test:"
curl -v http://${SERVER}/api/admin/user -H "Authorization: Bearer ${TOKEN}"
echo ""
