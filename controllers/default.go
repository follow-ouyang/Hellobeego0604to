package controllers

import (
	"HelloBeego0604to/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}


//默认请求为GET，如果是默认请求
func (c *MainController) Get() {
	// nameget := c.GetString("name")
	// ageget,err := c.GetInt("age")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

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
	//fmt.Println("post类型的请求")
	//user := c.Ctx.Request.FormValue("user")
	//fmt.Println("用户名为：",user)
	//psd := c.Ctx.Request.FormValue("psd")
	//fmt.Println("密码为:",psd)
	//
	////与固定值比较   用户名为：admin，密码为：123456
	//if user != "admin" || psd != "123456" {
	//	//失败页面
	//	c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
	//	return
	//}
	//	c.Ctx.ResponseWriter.Write([]byte("恭喜你，登陆成功"))

	dataByes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	var person models.Person
	err = json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名：",person.User,"年龄：",person.Age)
	c.Ctx.WriteString("用户名是："+person.User)



}