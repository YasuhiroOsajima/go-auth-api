#!/bin/sh

SERVER="192.168.33.13:9999"

curl -v -X POST http://${SERVER}/api/register --data '{"username":"test", "password":"passwd"}'
echo ""
