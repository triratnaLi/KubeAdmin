package main

import (
	_ "Ayile/routers"
	"Ayile/utils"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	utils.InitMysql()
	beego.Run()

}
