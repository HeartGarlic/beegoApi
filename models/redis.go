package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/cache"
)

type Redis struct {

}

// CACHE 初始化全局的缓存组件
var CACHE cache.Cache

func init()  {
	// 初始化redis缓存组件
	var err error
	CACHE, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil{
		fmt.Println(err.Error())
	}
}
