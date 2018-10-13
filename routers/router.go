package routers

import (
	"ybpro/controllers"
	"ybpro/controllers/test_controller"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})	// 默认使用GET方法
    beego.Router("/user", &controllers.MainController{})	// 默认使用GET方法
    beego.Router("/user/hello", &controllers.MainController{},"get:Hello")	// 默认使用GET方法
    // 测试
    beego.Router("/v1/test", &test_controller.TestController{}, "get:Test")
    beego.Router("/v1/test/:id:int", &test_controller.TestController{}, "get:GetTestNumber") // 正则表达
}
// 一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 8080 作为对外服务，另一个 8088 端口实行对内监控。