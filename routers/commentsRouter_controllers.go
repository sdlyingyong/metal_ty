package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["metal_ty/controllers:PortalController"] = append(beego.GlobalControllerRouter["metal_ty/controllers:PortalController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["metal_ty/controllers:PortalController"] = append(beego.GlobalControllerRouter["metal_ty/controllers:PortalController"],
        beego.ControllerComments{
            Method: "Article",
            Router: "/article/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
