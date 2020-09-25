package main

import (
	_ "HelloBeego0604to/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//定义config变量，接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	//config.string解析字符串类型
	appname := config.String("appname")
	fmt.Println("项目应用名称",appname)
	//config.int解析整数类型。。。。。
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		//panic：恐慌，程序恐慌
		panic("项目配置信息解析错误，请查验后重试")
	}
	fmt.Println("应用监听端口：",port)

	driver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbpassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")

	db,err := sql.Open(driver,dbUser+":"+dbpassword+"@tcp("+dbIP+")/"+dbName)
	if err != nil {
		panic("数据连接失败，请重试")
	}
	fmt.Println(db)

	beego.Run()
}

