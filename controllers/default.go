package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


//默认请求为GET，如果是默认请求
func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1298357346@qq.com"
	c.TplName = "index.tpl"
}
