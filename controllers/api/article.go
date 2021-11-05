package api

import (
	"github.com/beego/beego/v2/core/logs"
	"log"
	"metal_ty/controllers"
	. "metal_ty/models"
	"metal_ty/service"
	"strconv"
	"time"
)

type ArticleAPIController struct {
	controllers.AdminBaseController
}

/*
*  create one article
 */
func (c *ArticleAPIController) CreateArticle() {
	//handle exception return error data
	defer func() {
		if err := recover(); err != nil {
			c.Data["json"] = controllers.ErrorData(err.(error))
		}
		c.ServeJSON()
	}()

	//params
	title := c.GetString("title")
	if title == "" {
		log.Panic("title can not be empty")
	}
	content := c.GetString("content")
	if content == "" {
		log.Panic("content can not be empty")
	}
	category := c.GetString("category")
	keywords := c.GetString("keywords")

	//save db
	article := new(Article)
	article.Title = title
	article.Content = content
	article.Category = category
	article.Keywords = keywords
	article.Status = 1
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	articleService := service.NewService()
	_, err := articleService.Save(article)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		c.Data["json"] = controllers.SuccessData(nil)
	}

}

/**
 * get all article
 */
func (c *ArticleAPIController) ArticlesList() {
	//params
	args := c.GetString("search")
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	article := new(Article)
	param := map[string]string{
		"title": args,
	}
	list, total, err := article.GetArticlesByCondition(param, start, perPage)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = controllers.ErrorData(err)
	} else {
		data := map[string]interface{}{
			"result": list,
			"total":  total,
		}
		c.Data["json"] = controllers.SuccessData(data)
	}
	//return result
	c.ServeJSON()
}

/**
*	edit one article
 */
func (c *ArticleAPIController) ArticlesEdit() {
	//handle exception return error msg
	defer func() {
		if err := recover(); err != nil {
			c.Data["json"] = controllers.ErrorData(err.(error))
		}
		c.ServeJSON()
	}()

	//param
	artId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title := c.GetString("title")
	content := c.GetString("content")
	category := c.GetString("category")
	keywords := c.GetString("keywords")
	//save
	article := new(Article)
	article.Id = artId
	article.Title = title
	article.Content = content
	article.Category = category
	article.Keywords = keywords
	article.UpdatedAt = time.Now()
	_, err := article.Update()
	if err != nil {
		c.Data["json"] = controllers.ErrorData(err.(error))
		c.ServeJSON()
		return
	}
	c.Data["json"] = controllers.SuccessData(nil)
	c.ServeJSON()
}

/**
* delete one article
 */
func (c *ArticleAPIController) ArticlesDelete() {
	article := new(Article)
	artId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article.Id = artId
	article.Delete()
	c.Data["json"] = controllers.SuccessData(nil)
	c.ServeJSON()
}
