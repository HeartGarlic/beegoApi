package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// Category 分类管理
type Category struct {
	Id          int    `orm:"column(id)" json:"id"`
	Pid         int    `orm:"column(pid)" json:"pid"`
	Type        string `orm:"column(type)" json:"type"`
	Name        string `orm:"column(name)" json:"name"`
	Nickname    string `orm:"column(nickname)" json:"nickname"`
	Flag        string `orm:"column(flag)" json:"flag"`
	Image       string `orm:"column(image)" json:"image"`
	Keywords    string `orm:"column(keywords)" json:"keywords"`
	Description string `orm:"column(description)" json:"description"`
	DiyName     string `orm:"column(diyname)" json:"diyname"`
	CreateTime  int64  `orm:"column(createtime)" json:"createtime"`
	UpdateTime  int64  `orm:"column(updatetime)" json:"updatetime"`
	Weigh       int    `orm:"column(weigh)" json:"weigh"`
	Status      string `orm:"column(status)" json:"status"`
}

// TableName 表名
func (c *Category) TableName() string{
	return "fa_category"
}

// NewCategory 实例化方法
func NewCategory() *Category {
	return &Category{}
}

// AddOrUpdate 新增分类 / 修改分类
func (c *Category) AddOrUpdate() (id int, err error){
	// 初始化修改时间
	c.UpdateTime = time.Now().Unix()
	// 获取数连接句柄
	o := orm.NewOrm()
	if c.Id > 0 {
		// 查询数据
		err := o.Read(&Category{Id: c.Id})
		if err != nil {
			return 0, err
		}
		// 修改数据
		_, err = o.Update(c)
		if err != nil {
			return 0, err
		}
		return c.Id, nil
	}
	// 开始处理新增
	_, err = o.Insert(c)
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}

// DeleteCategory 删除分类
func (c *Category) DeleteCategory() error {
	o   := orm.NewOrm()
	err := o.Read(&Category{Id: c.Id})
	if err != nil {
		return err
	}
	_, err = o.Delete(&Category{Id: c.Id})
	if err != nil {
		return err
	}
	return nil
}

// GetCategory 获取分类
func (c *Category) GetCategory(params map[string]interface{}) (category []*Category, err error) {
	o  := orm.NewOrm()
	qs := o.QueryTable(c.TableName())
	// 拼接查询条件
	if params != nil {
		for k, v := range params {
			if v != "" {
				if k == "page" {
					// 计算分页
					qs = qs.Limit(10).Offset((v.(int) - 1) * 10)
				}else{
					qs = qs.Filter(k, v)
				}
			}
		}
	}
	_, err = qs.All(&category)
	if err != nil {
		return nil, err
	}
	return category, nil
}


