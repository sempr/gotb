package main

import (
	"fmt"
	"os"

	"github.com/sempr/gotb"
)

func main() {
	domain := os.Getenv("DOMAIN")
	appkey := os.Getenv("APPKEY")
	appsec := os.Getenv("APPSEC")
	cli := gotb.TopClient{}
	cli.Init(domain, appkey, appsec)
	r1, e1 := cli.Call("taobao.time.get", "", &gotb.TopParams{})
	fmt.Printf("%#v %#v\n", r1, e1)
	r2, e2 := cli.Call("taobao.shop.get", "", &gotb.TopParams{"nick": "朵朵云", "fields": "sid,cid,title,nick,pic_path,created,modified"})
	fmt.Printf("%#v %#v\n", r2, e2)
	r3, e3 := cli.Call("alibaba.geoip.get", "", &gotb.TopParams{"ip": "8.8.8.8", "language": "en"})
	fmt.Printf("%#v %#v\n", r3, e3)
}
