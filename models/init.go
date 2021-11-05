package models

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost, _ := beego.AppConfig.String("db.host")
	dbuser, _ := beego.AppConfig.String("db.user")
	dbpassword, _ := beego.AppConfig.String("db.password")
	dbport, _ := beego.AppConfig.String("db.port")
	dbname, _ := beego.AppConfig.String("db.name")
	//dbtimezone,_ := beego.AppConfig.String("db.timezone")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn)

	//orm.Debug = true
}
