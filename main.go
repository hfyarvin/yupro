package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "ybpro/routers" // 只引入执行了里面的 init 函数，
)

func main() {
	if beego.BConfig.RunMode == beego.DEV {
		orm.Debug = true
		beego.BConfig.Log.AccessLogs = true
	} else {
		beego.SetLevel(beego.LevelWarning) //default Trace
		beego.BConfig.Log.AccessLogs = false
	}
	// beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	// beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
	// 静态文件夹
	beego.BConfig.WebConfig.StaticDir["/"] = "static"
	beego.SetStaticPath("/down1", "download1")
	// beego.SetStaticPath("/images","images")
	// beego.SetStaticPath("/css","css")
	// beego.SetStaticPath("/js","js")

	beego.Run()
}

// func TransparentStatic(ctx *context.Context) {
// 	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
// 		return
// 	}
// 	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
// }
