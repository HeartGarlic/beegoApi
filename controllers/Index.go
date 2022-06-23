package controllers

import (
	"beegoApi/models"
	"github.com/beego/beego/v2/core/validation"
	"strconv"
)

type IndexController struct {
	BaseController
}

// Register 注册
func (c *IndexController) Register() {
	// 获取model
	user := models.NewUser()
	// 绑定表单数据到model
	err := c.ParseForm(user)
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return
	}
	// 开始数据验证
	valid := validation.Validation{}
	valid.Required(user.Username, "username").Message("用户名不能为空")
	valid.MinSize(user.Username, 2, "username").Message("用户名长度不能小于2")
	valid.MaxSize(user.Username, 20, "username").Message("用户名长度不能大于20")
	valid.Required(user.Nickname, "nickname").Message("昵称不能为空")
	valid.MinSize(user.Nickname, 2, "nickname").Message("昵称长度不能小于2")
	valid.MaxSize(user.Nickname, 20, "nickname").Message("昵称长度不能大于20")
	valid.Required(user.Password, "password").Message("密码不能为空")
	valid.MinSize(user.Password, 6, "password").Message("密码长度不能小于6")
	valid.MaxSize(user.Password, 20, "password").Message("密码长度不能大于20")
	valid.Required(user.Email, "email").Message("邮箱不能为空")
	valid.Email(user.Email, "email").Message("邮箱格式不正确")
	valid.Required(user.Mobile, "mobile").Message("手机号不能为空")
	valid.Mobile(user.Mobile, "mobile").Message("手机号格式不正确")
	valid.Required(user.Avatar, "avatar").Message("头像不能为空")
	// 开始验证
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(err.Message, nil, 400)
			return // 停止执行
		}
	}

	register, err := user.Register()
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return // 停止执行
	}
	c.Response("success", register, 200)
	return // 停止执行
}

// Login 登陆
func (c *IndexController) Login() {
	type loginForm struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	// 获取表单数据
	loginParams := &loginForm{}
	err := c.ParseForm(loginParams)
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return // 停止执行
	}
	// 开始数据验证
	valid := validation.Validation{}
	valid.Required(loginParams.Username, "username").Message("用户名不能为空")
	valid.MinSize(loginParams.Username, 2, "username").Message("用户名长度不能小于2")
	valid.MaxSize(loginParams.Username, 20, "username").Message("用户名长度不能大于20")
	valid.Required(loginParams.Password, "password").Message("密码不能为空")
	valid.MinSize(loginParams.Password, 6, "password").Message("密码长度不能小于6")
	valid.MaxSize(loginParams.Password, 20, "password").Message("密码长度不能大于20")
	// 开始验证
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(err.Message, nil, 400)
			return // 停止执行
		}
	}
	// 开始处理登陆流程
	userInfo, err := models.NewUser().Login(loginParams.Username, loginParams.Password)
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return
	}
	// 限制只返回部分字段
	response := make(map[string]interface{})
	response["id"] = userInfo.ID
	response["username"] = userInfo.Username
	response["nickname"] = userInfo.Nickname
	response["token"], _ = userInfo.GenerateToken()
	c.Response("success", response, 200)
}

// UserInfo 用户个人中心接口
func (c *IndexController) UserInfo() {
	userInfo := c.GetUser()
	c.Response("success", userInfo, 200)
}

// Upload 文件上传
func (c *IndexController) Upload(){
	// 判断文件类型? 判断文件大小?
	_, m, err := c.GetFile("file")
	if err != nil {
		c.Response(err.Error(), nil, 401)
		return
	}
	// 保存之前裁剪图片? 添加图片水印?
	// 保存文件
	err = c.SaveToFile("file", "public/upload/" + m.Filename)
	if err != nil {
		c.Response(err.Error(), nil, 401)
		return
	}
	returnMap := map[string]interface{}{
		"fileUrl": c.Ctx.Input.Scheme() + "://" + c.Ctx.Input.Host() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/public/upload/" + m.Filename,
	}
	c.Response("上传完成", returnMap, 200)
}
