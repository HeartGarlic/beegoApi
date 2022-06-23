// Package routers
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegoApi/controllers"
	"beegoApi/filter"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 尝试使用filter来实现路由权限控制
	// 注册v1版本命名空间
	ns := beego.NewNamespace("/v1",
		// 注册
		beego.NSRouter("/register", &controllers.IndexController{}, "post:Register"),
		// 登陆
		beego.NSRouter("/login", &controllers.IndexController{}, "post:Login"),
	)
	// 添加命名空间
	beego.AddNamespace(ns)
	authNs := beego.NewNamespace("/v2",
		// 授权信息校验
		beego.NSBefore(filter.AuthFilter),
		// 获取用户基础信息
		beego.NSRouter("/userinfo", &controllers.IndexController{}, "post:UserInfo"),
		// 测试文件上传
		beego.NSRouter("/upload", &controllers.IndexController{}, "post:Upload"),
	)
	// 添加命名空间
	beego.AddNamespace(authNs)
	// 添加或者修改分类
	beego.Router("/add-category", &controllers.CategoryController{}, "post:AddOrUpdate")
	// 删除分类
	beego.Router("/delete-category", &controllers.CategoryController{}, "get:Delete")
	// 分类列表
	beego.Router("/category-list", &controllers.CategoryController{}, "get:List")
}
