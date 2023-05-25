package routers

import (
	"Ayile/controllers"
	"Ayile/filters"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.InsertFilter("/resource/*", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/resource/query", &controllers.ResourceController{}, "get:Query;post:Add")

	//过滤路由
	beego.Router("/", &controllers.HomeController{}, "get:Get;post:Post")
	beego.InsertFilter("/", beego.BeforeRouter, filters.FilterUser)

	beego.Router("/logout", &controllers.LogoutController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/deployment", &controllers.DeploymentController{}, "get:Get;post:Post")

	beego.Router("/daemonset", &controllers.DaemonSetController{}, "get:Get;post:Post")
}
