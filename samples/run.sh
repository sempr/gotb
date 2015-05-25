#!/bin/bash
go build main.go
DOMAIN="gw.api.taobao.com" APPKEY="Your APP KEY" APPSEC="Your APP SECRET" ./main
