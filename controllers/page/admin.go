package page

import (
	"github.com/beego/beego/v2/core/logs"
	"metal_ty/controllers"
	."metal_ty/models"
)

type AdminPageController struct {
	controllers.AdminBaseController
}


//user module
func(c *AdminPageController) Welcome () {
	c.TplName = "admin/index.html"
}


//admin module


//article  page
func(c *AdminPageController) CreateArticle() {
	c.TplName = "admin/article-add.html"
}

func (c *AdminPageController) ArticleList () {
	c.TplName = "admin/article-list.html"
}

func (c *AdminPageController) ArticleEdit () {
	//get article detail
	article := new(Article)
	artId,_ := c.GetInt("id")
	article.Id = int(artId)
	article.GetById()
	logs.Debug(article)

	//display list
	c.Data["article"] = article
	c.TplName = "admin/article-edit.html"
}