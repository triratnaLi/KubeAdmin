package base

import (
	"Ayile/models"
	"Ayile/utils"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

}

func (this *LoginController) Post() {

	err := this.Ctx.Request.ParseForm()
	if err != nil {
		return
	}
	username := this.GetString("username")
	password := this.GetString("password")

	user := models.QueryUserWithUsername(username)
	if user == 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
		this.ServeJSON()
		return
	}

	// 验证密码是否正确
	passwordMD5 := utils.MD5(password)
	if passwordMD5 != utils.MD5(password) {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}

	this.SetSession("user", username)
	this.Redirect("/", 302)

}
