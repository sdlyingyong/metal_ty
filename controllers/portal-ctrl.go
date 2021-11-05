package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "metal_ty/models"
	"metal_ty/util"
	"regexp"
	"strconv"
)

//register controller
type PortalController struct {
	beego.Controller
}

func (this *PortalController) Prepare() {
	val, _ := beego.AppConfig.String("runmode")
	this.Data["env"] = val
}

//home page
// @router / [get]
func (c *PortalController) Get() {

	//param
	pageNo, err := c.GetInt("pageNo")
	if err != nil {
		pageNo = 1
	}
	pageSize, err := c.GetInt("pageSize")
	if err != nil {
		pageSize = 10
	}
	skip := (pageNo - 1) * pageSize

	//get data
	params := map[string]string{}
	article := &Article{}
	articleList, total, err := article.GetArticlesByCondition(params, skip, pageSize)
	if err != nil {
		//empty data
		c.Data["articleList"] = []Article{}
		c.Data["total"] = 0
	} else {
		//dynamic array
		var artList = make([]ArticlePortal, len(articleList))
		//json data style
		for index, art := range articleList {
			//html tags
			re := regexp.MustCompile("\\<[\\S\\s]+?\\>")
			//img tags
			reimg := regexp.MustCompile(`<img (\S*?)[^>]*>.*?|<.*? />`)
			htmlStr := util.Md2html(art.Content)
			artList[index].Id = art.Id
			artList[index].Title = art.Title
			artList[index].Content = beego.Substr(re.ReplaceAllString(htmlStr, ""), 0, 30)
			artList[index].Img = string(reimg.Find([]byte(htmlStr)))
			artList[index].Status = art.Status
			artList[index].Category = art.Category
			artList[index].ViewCount, err = article.GetArticleViewCount(art.Id)
			artList[index].UpdatedAt = art.UpdatedAt.Format("2016-01-02")
		}
		c.Data["articleList"] = artList
		c.Data["total"] = total
		c.Data["pageNo"] = pageNo
		c.Data["pageSize"] = pageSize
	}
	//request ip addr
	logs.Info("访问ip:", c.Ctx.Input.Header("Remote_addr"))
	//display page
	c.TplName = "index.html"
}

//get one article
//@router /article/:id [get]
func (c *PortalController) Article() {
	//param
	artId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Data["content"] = err
		c.TplName = "404.html"
	} else {
		//get one article data
		article := &Article{}
		article.Id = int(artId)
		articlePortal, err := article.ArticleDetail()

		logs.Info("article portal is : ")
		logs.Info(articlePortal)

		//todo save article log

		if err != nil {
			logs.Error(err)
		}
		articlePortal.Title = article.Title
		articlePortal.Keywords = article.Keywords
		articlePortal.Content = article.Content
		articlePortal.Category = article.Category
		articlePortal.UpdatedAt = article.UpdatedAt.Format("2006-01-02 15:04")
		c.Data["article"] = articlePortal
		c.Data["zero"] = int(0)
		//display
		c.TplName = "article.html"

	}

}
