package base

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
	//清除该用户登录状态的数据
	this.DelSession("user")

	this.Redirect("/login", 302)
}
