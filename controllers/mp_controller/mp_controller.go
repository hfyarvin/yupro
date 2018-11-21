package mp_controller

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/util"
	"ybpro/libs/wx_mp"
)

type MpController struct {
	beego.Controller
}

/*
* appid: wx72a01fd5ef9a40d8
* appsecret: 6d3f12f767ad8836d01dde057d6f476f
* url: http://mpdrg9.natappfree.cc/v1/mp/wx72a01fd5ef9a40d8
* token: arvinWongWxToken
 */


//测试地址http://tgo.xikego.com:28888/v1/mp/wxb85becdc77d227c1?timestamp=1452007620&nonce=1530712726&signature=4b196dfa7023b75cdebd8a9d2dbccfff601bae7e&echostr=1479928067271949302
//微信公众平台验证url和token的有效性
func (self *MpController) Handler() {
	beego.BeeLogger.Warn("**************wx service*******************")
	t := self.GetString("timestamp")
	n := self.GetString("nonce")
	s := self.GetString("signature")
	e := self.GetString("echostr")
	mpappid := self.Ctx.Input.Param(":mpappid")
	if mpappid == "" {
		m := "mpappid is null"
		beego.BeeLogger.Error(m)
		self.Ctx.WriteString(m)
		return
	}
	// gzh := wx_gongzhonghao_model.GetGZHByAppId(mpappid, false)
	// token := ""
	if mpappid != "wx72a01fd5ef9a40d8" {
		m := "mpappid is not allowed"
		beego.BeeLogger.Error(m)
		self.Ctx.WriteString(m)
		return
	}
	// token = "arvinWongWxToken"
	gzh := new(wx_mp.WxGzh)
	gzh.AppId = "wx72a01fd5ef9a40d8"
	gzh.Token = "arvinWongWxToken"
	gzh.AppSecret = "6d3f12f767ad8836d01dde057d6f476f"
	//在开发者首次提交验证申请时，微信服务器将发送GET请求到填写的URL上，
	//并且带上四个参数（signature、timestamp、nonce、echostr），
	//开发者通过对签名（即signature）的效验，来判断此条消息的真实性。
	if self.Ctx.Request.Method == "GET" {
		if t == "" || n == "" || s == "" || e == "" {
			info := fmt.Sprintf("params,timestamp=[%s],nonce=[%s],signature=[%s],echostr=[%s]", t, n, s, e)
			beego.BeeLogger.Error(info)
			self.Ctx.WriteString(info)
			return
		}
		// 验证签名
		//密文模式
		// sign := util.MsgSign(gzh.Token, t, n, encryptedMsg)
		// 明文模式
		sign := util.Sign(gzh.Token, t, n)
		if sign != s {
			beego.BeeLogger.Warn("wx authitication sign:%s, s:%s", sign, s)
			text := fmt.Sprintf("check signature fail,sign=[%s],params[sign]=[%s]", sign, s)
			self.Ctx.WriteString(text)
			return
		}
		beego.BeeLogger.Warn("params,timestamp=[%s],nonce=[%s],signature=[%s],echostr=[%s]", t, n, s, e)
		self.Ctx.WriteString(e)
	} else { //处理其他微信请求
		if err := wx_mp.HandleWithRequest(self.Ctx, gzh); err != nil {
			beego.BeeLogger.Error("Handle with request error[%s]", err)
		}
		self.Ctx.WriteString("success")
		return
	}

}
