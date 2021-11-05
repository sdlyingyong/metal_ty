package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"metal_ty/controllers"
	"metal_ty/controllers/api"
	"metal_ty/controllers/page"
)

func init() {
	//all endpoint

	//portal
	beego.Include(
		//user web endpoint
		&controllers.PortalController{},
	)

	//admin module
	ns := beego.NewNamespace("/admin",
		beego.NSRouter("/", &page.AdminPageController{}, "get:Welcome"),
		beego.NSNamespace("/api",
			//article crud
			beego.NSRouter("/article", &api.ArticleAPIController{}, "post:CreateArticle"),
			beego.NSRouter("/articles", &api.ArticleAPIController{}, "get:ArticlesList"),
			beego.NSRouter("/article/:id", &api.ArticleAPIController{}, "put:ArticlesEdit"),
			beego.NSRouter("/article/:id", &api.ArticleAPIController{}, "delete:ArticlesDelete"),
		),
		beego.NSNamespace("/page",
			//home page
			beego.NSRouter("/", &page.AdminPageController{}, "get:Welcome"),
			//article show page
			beego.NSRouter("/article-add", &page.AdminPageController{}, "get:CreateArticle"),
			beego.NSRouter("/article-list", &page.AdminPageController{}, "get:ArticleList"),
			beego.NSRouter("/article-edit", &page.AdminPageController{}, "get:ArticleEdit"),
		),
	)

	beego.AddNamespace(ns)

}
