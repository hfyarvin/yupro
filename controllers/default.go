package controllers

import (
	// "fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller	// 继承
}

func (c *MainController) Get() {	//重写GET方法
	beego.BeeLogger.Warn("get ********")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl" // 如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找
}

func (c *MainController) Hello()  {
	beego.BeeLogger.Warn("this is a log****************")
	c.Ctx.WriteString("Hello")	// 直接输出字符串
}