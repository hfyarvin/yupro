package test_controller

import (
	"fmt"
)

func (self *TestController) Test()  {
	ret := "sss"
	self.Data["json"] = ret
	self.ServeJSON()
}

func (self *TestController) GetTestNumber()  {
	_id := self.Ctx.Input.Param(":id")
	ret := fmt.Sprintf("the number of test: %s", _id)
	self.Data["json"] = ret
	self.ServeJSON()
}