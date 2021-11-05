package main

import (
	"metal_ty/models"
	_ "metal_ty/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	models.Init()
	beego.Run()
}

