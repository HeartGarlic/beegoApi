package filter

import (
	"beegoApi/models"
	"github.com/beego/beego/v2/server/web/context"
)

// AuthorizationKey header 中的 token 名称
const AuthorizationKey = "Authorization"

func AuthFilter (ctx *context.Context) {
	// 校验 token 从 header 中校验
	// 校验成功后，将用户信息放入 context 中
	// 校验失败，返回 401
	authToken := ctx.Input.Header(AuthorizationKey)
	if authToken == "" {
		ctx.Output.SetStatus(401)
		err := ctx.Output.Body([]byte("Unauthorized is empty"))
		if err != nil {
			return
		}
		return
	}
	// 开始解析token
	decrypt, err := models.NewUser().ParseToken(authToken)
	if err != nil {
		ctx.Output.SetStatus(401)
		err := ctx.Output.Body([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	if decrypt == ""{
		ctx.Output.SetStatus(401)
		err := ctx.Output.Body([]byte("Unauthorized is empty"))
		if err != nil {
			return
		}
		return
	}
}