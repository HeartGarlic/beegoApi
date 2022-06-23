package models

import (
	"github.com/beego/beego/v2/client/orm"
)
import _ "github.com/go-sql-driver/mysql"

func init() {
	// 开始sql调试
	orm.Debug = true
	// 注册models
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(User))
	// 注册数据库
	_ = orm.RegisterDataBase("default", "mysql", "")
}
