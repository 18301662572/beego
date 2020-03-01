package sysint

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//数据库初始化

func initDB() {
	//连接名称
	dbAlias := beego.AppConfig.String("db_alias")
	//数据库名称
	dbName := beego.AppConfig.String("db_name")
	//数据库连接用户
	dbUser := beego.AppConfig.String("db_user")
	//数据库连接密码
	dbPwd := beego.AppConfig.String("db_pwd")
	//数据库IP(域名)
	dbHost := beego.AppConfig.String("db_host")
	//数据库端口 (127.0.0.1:通过网卡，有防火墙限制； localhost:不经过网卡，没有防火墙限制)
	dbPort := beego.AppConfig.String("db_port")
	//charset
	dbCharset := beego.AppConfig.String("db_charset")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp"+"("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset)
	fmt.Println(dbUser + ":" + dbPwd + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset)

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动创建表
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}
