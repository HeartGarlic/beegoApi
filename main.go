package main

import (
	_ "beegoApi/models"
	_ "beegoApi/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// 注册静态文件访问地址
	go beego.SetStaticPath("/public","public")

	beego.Run()
}
