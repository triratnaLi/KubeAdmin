package utils

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	//切记：导入驱动包
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("开始初始化mysql")

	driverName, _ := beego.AppConfig.String("driverName")
	user, _ := beego.AppConfig.String("mysqluser")
	pwd, _ := beego.AppConfig.String("mysqlpwd")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	dbname, _ := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/yile?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		//创建用户表
		CreateTableWithUser()
		//创建文章表
		CreateTableWithArticle()

	}
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

// 创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

// 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
