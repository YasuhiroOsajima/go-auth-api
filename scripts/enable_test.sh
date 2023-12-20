#!/bin/sh

SERVER="192.168.33.13:9999"
COOKIE_FILE="/tmp/$(date +'%Y%m%d%H%M%S').cookie"

curl -c "${COOKIE_FILE}" --silent -X POST http://${SERVER}/api/login --data '{"username": "test", "password": "passwd"}'
echo ""
cat "${COOKIE_FILE}"
echo ""

curl -v -X POST http://${SERVER}/api/admin/enable -b "${COOKIE_FILE}" --data '{"username":"test"}'
echo ""

rm -f "${COOKIE_FILE}"
