package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


//默认请求为GET，如果是默认请求
func (c *MainController) Get() {
	//获取get类型请求的数据
	name :=  c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	//以admin，19为正确数据
	if name != "admin" || age != "19" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误……"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功^_^"))

	c.Data["Website"] = "https://home.firefoxchina.cn/"
	c.Data["Email"] = "1298357346@qq.com"
	c.TplName = "index.tpl"
}

/**
	该post方法是处理post类型的请求的时候要调用的方法
*/
func (c *MainController) Post() {
	fmt.Println("post类型的请求")
	user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码为:",psd)
	
	//与固定值比较   用户名为：admin，密码为：123456
	if user != "admin" || psd != "123456" {
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}
		c.Ctx.ResponseWriter.Write([]byte("恭喜你，登陆成功"))
}