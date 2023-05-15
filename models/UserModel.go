package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         int    `orm:"column(id)"`
	Username   string `orm:"column(username)"`
	Password   string `orm:"column(password)"`
	Status     int    `orm:"column(status)"` // 0 正常状态， 1删除
	Createtime int64  `orm:"column(createtime);type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/yile?charset=utf8mb4")

	// 自动建表
	orm.RunSyncdb("default", false, true)

}

func InsertUser(user User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	return id, err
}

// 按条件查询
func QueryUserWightCon(con string) int {
	o := orm.NewOrm()
	var user User
	o.Raw("select id from user " + con).QueryRow(&user)
	return user.Id
}

// 根据用户名查询id
// func QueryUserWithUsername(username string) (*User, error) {
func QueryUserWithUsername(username string) int {
	o := orm.NewOrm()
	user := new(User)
	err := o.QueryTable(user).Filter("username", username).One(user)
	if err != nil {
		return 0
	}
	return user.Id
}

// 根据用户名和密码，查询id
func QueryUserWithParam(username, password string) (*User, error) {
	o := orm.NewOrm()
	user := new(User)
	err := o.QueryTable(user).Filter("username", username).Filter("password", password).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
