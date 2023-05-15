package routers

import (
	"Ayile/controllers"
	"Ayile/filters"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//beego.InsertFilter("/", beego.BeforeRouter, filters.FilterUser)

	beego.Router("/", &controllers.HomeController{}, "get:Get;post:Post")
	beego.InsertFilter("/", beego.BeforeRouter, filters.FilterUser)

	beego.Router("/logout", &controllers.LogoutController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})

	//部署概览
	beego.Router("/depupload", &controllers.DepUploadController{})
	//前端部署
	beego.Router("/frontend", &controllers.FrontendController{})

}
