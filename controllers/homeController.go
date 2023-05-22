package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

/**
 * 请求：http://localhost:8080/
 * 请求类型：Get
 * 请求描述：
 */

func (this *HomeController) Get() {
	this.Data["User"] = this.GetSession("user")

	currentPath := this.Ctx.Request.URL.Path
	this.Data["Page"] = currentPath
	this.TplName = "home.html"
}
