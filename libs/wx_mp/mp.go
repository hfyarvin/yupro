package wx_mp

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/message/response"
	// "github.com/chanxuehong/wechat/mp/jssdk"
)

type WxGzh struct {
	Token string
	AppId string
	AppSecret string
}


// 处理微信消息被动消息入口
func HandleWithRequest(ctx *context.Context, gzh *WxGzh) error {
	r := ctx.Request // 请求
	resp := ctx.ResponseWriter

	// 明文
	rawMsgXML, err := ioutil.ReadAll(r.Body) //读取XML
	if err != nil {
		beego.BeeLogger.Error("ioutil.ReadAll err=[%s],body=[%v]", err, r.Body)
		return err
	}
	body := rawMsgXML

	// 解析至mp.MixedMessage
	mixedMessage := new(mp.MixedMessage)
	if err := xml.Unmarshal([]byte(body), mixedMessage); err != nil {
		beego.BeeLogger.Error("HandleWithRequest decode xml body err=[%s]", err)
		return err
	}
	// 更新会话时间TODO

	// 信息类型
	msgType := mixedMessage.MsgType
	server := mixedMessage.ToUserName
	openid := mixedMessage.FromUserName
	beego.BeeLogger.Warn("msg_type: %s, to_user: %s, from_user: %s.", msgType, server, openid)

	if openid == "xxxxx" {
		userInfo := UserInfo(openid, gzh.AppId, gzh.AppSecret)
		beego.BeeLogger.Warn(fmt.Sprintf("%v", userInfo))
	}

	switch msgType {
	case "image": 
	return HandleWithImage(server, openid, mixedMessage, gzh, resp)
	default: 
	return HandleWithDefault(server, openid, mixedMessage, gzh, resp)
	}
	return nil
}

// 默认处理方式
func HandleWithDefault(server, openid string, mixedMessage *mp.MixedMessage,
gzh *WxGzh, resp http.ResponseWriter) error {
	return nil
}

//处理图片消息
func HandleWithImage(server, openid string, mixedMessage *mp.MixedMessage,
	gzh *WxGzh, resp http.ResponseWriter) error {
	ret := sendText(server, openid, "图片已收到")
	xml.NewEncoder(resp).Encode(ret)
	return nil
}


// 返回文本消息
func sendText(server, openid, respContent string) (text *response.Text) {
	t := time.Now().Unix()
	return response.NewText(openid, server, t, respContent)
}

// get client
func GetClt(appId, appSecret string) *mp.Client { 
	// beego.BeeLogger.Error("wx_mp.getClt('%s') too many", appId)
	//获取access token
	srv := getAccessTokenServer(appId, appSecret)
	_, err := srv.Token()
	if err != nil {
		beego.BeeLogger.Error("access token err=[%s]", err)
	}
	// beego.BeeLogger.Error("wx_mp.getClt('%s') token='%s'", appId, token)
	cli := mp.NewClient(srv, nil)
	return cli

}


var lock sync.Mutex
// 获取accessTokenServer
func getAccessTokenServer(appId, appSecret string) *mp.DefaultAccessTokenServer {
	// k := fmt.Sprintf("getAccessTokenServerByAppId_%s", appId)
	// srv, ok := cache.Cache.Get(k).(*mp.DefaultAccessTokenServer)
	// if ok {
	// 	return srv
	// }
	lock.Lock()
	// 高并发下先前在等待的可能已经获取到了，直接返回
	// srv, ok = cache.Cache.Get(k).(*mp.DefaultAccessTokenServer)
	// if ok {
	// 	return srv
	// }
	srv := mp.NewDefaultAccessTokenServer(appId, appSecret, nil)
	beego.BeeLogger.Warn("wx_mp.getAccessTokenServer('%s') too many", appId)
	// if err := cache.Cache.Put(k, srv, 24*time.Hour); err != nil {
	// 	beego.BeeLogger.Error("New AccessTokenServer Error=[%s]", err)
	// }
	lock.Unlock()
	return srv
}