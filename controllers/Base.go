package controllers

import (
	"beegoApi/filter"
	"beegoApi/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}


// UserInfo 登陆解析token的用户信息
type UserInfo struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Mobile    string `json:"mobile"`
	Timestamp int64  `json:"timestamp"`
}

// GetUser 获取用户信息
func (b *BaseController) GetUser() *UserInfo {
	token := b.Ctx.Input.Header(filter.AuthorizationKey)
	if token == "" {
		b.Response("token is empty", nil, 401)
		return nil
	}
	parseToken, err := models.NewUser().ParseToken(token)
	if err != nil {
		b.Response(err.Error(), nil, 400)
		return nil
	}
	jsonModel := &UserInfo{}
	err = json.Unmarshal([]byte(parseToken), jsonModel)
	if err != nil {
		b.Response(err.Error(), nil, 400)
		return nil
	}
	return jsonModel
}

// Response 返回数据
func (b *BaseController) Response(msg string, data interface{}, code int) {
	b.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	_ = b.ServeJSON()
}
