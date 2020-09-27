package main

import (
	_ "HelloBeego0604to/routers"
	"github.com/astaxie/beego"
)

func main() {
	//定义config变量，接收并赋值为全局配置变量

	////获取配置选项
	////config.string解析字符串类型
	//appName := config.String("appname")
	//fmt.Println("项目应用名称:",appName)
	////config.int解析整数类型。。。。。
	//port,err := config.Int("httpport")
	//if err != nil {
	//	//配置信息解析错误
	//	//panic：恐慌，程序恐慌
	//	fmt.Println(err.Error())
	//	panic("项目配置信息解析错误，请查验后重试")
	//}
	//fmt.Println("应用监听端口：",port)

	//1、连接数据库
	//db_mysql.Connect()

	beego.Run()
}

