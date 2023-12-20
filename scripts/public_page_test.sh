#!/bin/sh

WEB_SERVER="192.168.33.13"

echo "Access public page patter test:"
curl -v http://${WEB_SERVER}/login/index.html
echo ""

echo ""
echo "Unauthorized patter test:"
curl -v http://${WEB_SERVER}/private/index.html
echo ""
