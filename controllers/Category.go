package controllers

import (
	"beegoApi/models"
	"github.com/beego/beego/v2/core/validation"
)

// CategoryController 分类控制器
type CategoryController struct {
	BaseController
}

// AddOrUpdate 新增分类
func (c *CategoryController) AddOrUpdate() {
	// 开始参数绑定验证
	type requestParams struct {
		Id          int    `form:"id"`
		Name        string `form:"name"`
		Type        string `form:"type"`
		Nickname    string `form:"nickname"`
		Flag        string `form:"flag"`
		Image       string `form:"image"`
		Keywords    string `form:"keywords"`
		Description string `form:"description"`
		Status      string `form:"status"`
	}
	// 开始参数绑定
	var params requestParams
	err := c.ParseForm(&params)
	if err != nil {
		c.Response(err.Error(), nil, 400)
	}
	// 开始参数校验
	valid := validation.Validation{}
	valid.Required(params.Name, "name").Message("名称不能为空")
	valid.Required(params.Type, "type").Message("类型不能为空")
	valid.Required(params.Nickname, "nickname").Message("别名不能为空")
	valid.Required(params.Flag, "flag").Message("标识不能为空")
	valid.Required(params.Image, "image").Message("图片不能为空")
	valid.Required(params.Keywords, "keywords").Message("关键字不能为空")
	valid.Required(params.Description, "description").Message("描述不能为空")
	valid.Required(params.Status, "status").Message("状态不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Response(err.Message, nil, 400)
			return
		}
	}
	// 调用模型新增或修改
	category := models.NewCategory()
	if params.Id > 0{
		category.Id = params.Id
	}
	category.Name = params.Name
	category.Type = params.Type
	category.Nickname = params.Nickname
	category.Flag = params.Flag
	category.Image = params.Image
	category.Keywords = params.Keywords
	category.Description = params.Description
	category.Status = params.Status
	update, err := category.AddOrUpdate()
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return
	}
	// 声明返回值map
	returnMap := map[string]interface{}{
		"update": update,
	}
	c.Response("success", returnMap, 200)
}

// Delete 删除分类
func (c *CategoryController) Delete() {
	id, _ := c.GetInt("id")
	if id <= 0 {
		c.Response("id不能为空", nil, 400)
		return
	}
	category   := models.NewCategory()
	category.Id = id
	err := category.DeleteCategory()
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return
	}
	// 定义返回值map
	returnMap := map[string]interface{}{
		"delete": true,
	}
	c.Response("success", returnMap, 200)
	return
}

// List 查询分类列表
func (c *CategoryController) List() {
	page, _  := c.GetInt("page", 0)
	// 获取请求参数
	params   := map[string]interface{}{
		"page": page,
		"name": c.GetString("name", ""),
	}
	category         := models.NewCategory()
	getCategory, err := category.GetCategory(params)
	if err != nil {
		c.Response(err.Error(), nil, 400)
		return
	}
	// 声明返回值map
	returnMap := map[string]interface{}{
		"list": getCategory,
	}
	c.Response("success", returnMap, 200)
	return
}
