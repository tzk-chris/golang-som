package main

import (
	_ "chrixblob/models"
	_ "chrixblob/routers"
	"chrixblob/tools"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// DB -- MySQL
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")
	data_source := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", data_source, 30)

	fmt.Printf("连接成功! \n%s", data_source)
	// 自动创建表格
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		panic(err)
	}

}

func main() {

	beego.InsertFilter("/cms/main/*", beego.BeforeRouter, tools.CmsLoginFilter)
	//beego.InsertFilter("/comment", beego.BeforeRouter, tools.FrontLoginFilter)  // 在后台已经定义好登录限制
	orm.RunCommand()
	beego.Run()
}
