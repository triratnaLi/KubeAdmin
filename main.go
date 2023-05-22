package main

import (
	_ "Ayile/routers"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	beego.Run()

}
