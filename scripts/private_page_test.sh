#!/bin/sh

SERVER="127.0.0.1:9999"
WEB_SERVER="127.0.0.1"
COOKIE_FILE="/tmp/$(date +'%Y%m%d%H%M%S').cookie"

curl -c "${COOKIE_FILE}" --silent -X POST http://${SERVER}/api/login --data '{"username": "test", "password": "passwd"}'
echo ""
cat "${COOKIE_FILE}"
echo ""

curl -v -b "${COOKIE_FILE}" http://${WEB_SERVER}/private/index.html
echo ""

rm -f "${COOKIE_FILE}"
