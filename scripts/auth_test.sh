#!/bin/sh

SERVER="127.0.0.1:9999"
COOKIE_FILE="/tmp/$(date +'%Y%m%d%H%M%S').cookie"

curl -c "${COOKIE_FILE}" --silent -X POST http://${SERVER}/api/login --data '{"username": "test", "password": "passwd"}'
echo ""
cat "${COOKIE_FILE}"
echo ""

echo "NG patter test:"
curl -v http://${SERVER}/api/admin/user
echo ""

echo ""
echo "OK pattern test:"
curl -v http://${SERVER}/api/admin/user -b "${COOKIE_FILE}"
echo ""

rm -f "${COOKIE_FILE}"
