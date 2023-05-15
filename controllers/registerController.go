package controllers

import (
	"Ayile/models"
	"Ayile/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"

}
func (this *RegisterController) Post() {
	//获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	if id := models.QueryUserWithUsername(username); id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已存在"}
		this.ServeJSON()
		return
	}

	//判断用户名和密码是否为空
	if username == "" || password == "" || repassword == "" {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "用户名或密码为空"}
		this.ServeJSON()
		return
	}

	// 注册用户名和密码
	// 存储的密码是 md5 后的数据，那么在登录验证的时候，也是需要将用户的密码 md5 之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	user := models.User{Username: username, Password: password, Createtime: time.Now().Unix()}
	if _, err := models.InsertUser(user); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}

	// 注册成功后重定向到首页
	this.Redirect("/login", 302)
}
