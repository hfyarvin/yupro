package main

import (
	_ "ybpro/routers"	// 只引入执行了里面的 init 函数，
	"github.com/astaxie/beego"
)

func main() {
	// 静态文件夹
	beego.SetStaticPath("/down1", "download1")
	beego.SetStaticPath("/images","images")
	beego.SetStaticPath("/css","css")
	beego.SetStaticPath("/js","js")
	beego.Run()
}

