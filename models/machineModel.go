package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
)

type Machine struct {
	Id         int    `orm:"column(Id)"`
	Type       string `orm:"column(Type)"`
	UseType    string `orm:"column(UseType)"`
	OsVersion  string `orm:"column(OsVersion)"` // 0 正常状态， 1删除
	Brand      string `orm:"column(Brand)"`     //品牌
	CPU        int    `orm:"column(CPU)"`
	Memory     int    `orm:"column(Memory)"`
	SystemDisk int    `orm:"column(SystemDisk)"`
	DataDisk   int    `orm:"column(DataDisk)"`
	IpAddress  string `orm:"column(IpAddress)"`
}

func init() {
	// 加载配置文件
	appConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	dbHost, _ := appConfig.String("host")
	dbPort, _ := appConfig.String("port")
	dbUser, _ := appConfig.String("mysqluser")
	dbPassword, _ := appConfig.String("mysqlpwd")
	dbName, _ := appConfig.String("dbname")

	dbConnStr := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	orm.RegisterDataBase("default", "mysql", dbConnStr)

	orm.RegisterModel(new(Machine))
	// 自动建表
	orm.RunSyncdb("default", false, true)

}
