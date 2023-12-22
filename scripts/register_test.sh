#!/bin/sh

SERVER="127.0.0.1:9999"

curl -v -X POST http://${SERVER}/api/register --data '{"username":"test", "password":"passwd"}'
echo ""
