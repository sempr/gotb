package main

import (
	"fmt"
	"os"

	"github.com/sempr/gotb"
)

func show(cli gotb.TopClient, method string, token string, args *gotb.TopParams) {
	r3, e3 := cli.Call(method, token, args)
	fmt.Printf("%#v %#v\n", r3, e3)
}

func main() {
	domain := os.Getenv("DOMAIN")
	appkey := os.Getenv("APPKEY")
	appsec := os.Getenv("APPSEC")
	cli := gotb.TopClient{}
	cli.Init(domain, appkey, appsec)
	show(cli, "taobao.time.get", "", &gotb.TopParams{})
	show(cli, "taobao.shop.get", "", &gotb.TopParams{"nick": "朵朵云", "fields": "sid,cid,title,nick,pic_path,created,modified"})
	show(cli, "alibaba.geoip.get", "", &gotb.TopParams{"ip": "8.8.8.8", "language": "en"})
	show(cli, "taobao.areas.get", "", &gotb.TopParams{"fields": "type,name,parent_id,zip,id"})
}
