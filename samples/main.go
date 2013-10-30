package main

import (
	"github.com/sempr/gotb"
)

func main() {
	cli := gotb.TopClient{}
	cli.Init("gw.api.taobao.com", "21660263", "829bf4c687e6777b2b03cf9de6abf55d")
	cli.Call("taobao.shop.get", "", &gotb.TopParams{"nick": "朵朵云", "fields": "sid,cid,title,nick,pic_path,created,modified"})
	cli.Call("taobao.item.get", "", &gotb.TopParams{"num_iid": "18705144228", "fields":"title,price,num_iid,cid,sku.price"})
}
