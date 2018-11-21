package wx_mp

import (
	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/mp/user"
)

// 获取用户信息
func UserInfo(openId, appId, appSecret string) *user.UserInfo {
	clt := GetClt(appId, appSecret)
	if info, err := (*user.Client)(clt).UserInfo(openId, user.Language_zh_CN); err != nil {
		beego.BeeLogger.Error("mp get user info.err=[%s],opengId=[%s]", err, openId)
		return nil
	} else {
		return info
	}
}