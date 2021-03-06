package routers

import (
	"github.com/astaxie/beego"
	"ybpro/controllers"
	"ybpro/controllers/test_controller"
	"ybpro/controllers/mp_controller"
)

func init() {
	beego.Router("/", &controllers.MainController{})                        // 默认使用GET方法
	beego.Router("/user", &controllers.MainController{})                    // 默认使用GET方法
	beego.Router("/user/hello", &controllers.MainController{}, "get:Hello") // 默认使用GET方法
	// 测试
	beego.Router("/v1/test", &test_controller.TestController{}, "get:Test")
	beego.Router("/v1/test/path", &test_controller.TestController{}, "get:GetPath")
	beego.Router("/v1/test/http", &test_controller.TestController{}, "get:GetRequestInfo")
	beego.Router("/v1/test/:id:int", &test_controller.TestController{}, "get:GetTestNumber") // 正则表达

	// 公众号
	beego.Router("/mp/:mpappid:string", &mp_controller.MpController{}, "*:Handler")
}

// 一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 8080 作为对外服务，另一个 8088 端口实行对内监控。
