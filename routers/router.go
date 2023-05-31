package routers

import (
	"Ayile/controllers/base"
	"Ayile/controllers/kubernetes"
	"Ayile/controllers/machine"
	"Ayile/filters"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//base
	beego.Router("/", &base.HomeController{}, "get:Get;post:Post")
	beego.InsertFilter("/", beego.BeforeRouter, filters.FilterUser)

	beego.Router("/logout", &base.LogoutController{})

	beego.Router("/login", &base.LoginController{})
	beego.Router("/register", &base.RegisterController{})

	//resoure
	beego.InsertFilter("/resource/*", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/resource/query", &machine.ResourceController{}, "get:Query;post:Add")

	//kubernetes
	beego.InsertFilter("/kube/*", beego.BeforeRouter, filters.FilterUser)
	beego.Router("/kube/deployment", &kubernetes.DeploymentController{}, "get:Get;post:Post")
	beego.Router("/kube/daemonset", &kubernetes.DaemonSetController{}, "get:QueryDaemonSets;post:Post")
	beego.Router("/kube/service", &kubernetes.ServicesController{}, "get:QueryServices;post:Post")
}
