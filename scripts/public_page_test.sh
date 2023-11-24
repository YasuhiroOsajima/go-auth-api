#!/bin/sh

echo "Access public page patter test:"
curl -v http://127.0.0.1/login/index.html
echo ""

echo ""
echo "Unauthorized patter test:"
curl -v http://127.0.0.1/private/index.html
echo ""
