package controllers

import (
	"HelloBeego0604to/db_mysql"
	"HelloBeego0604to/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//解析前端提交的用户注册数据
	dataBytes,err :=  ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(dataBytes,&user)
	if err != nil {
		r.Ctx.WriteString("数据解析错误，请重试")
		return
	}
	//一切正常，将用户信息保存到数据库中去
	//直接调用保存数据的函数，并判断保存后的结果
	row,err :=  db_mysql.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("注册用户信息失败，请重试")
		return
	}
	fmt.Println(row)
	r.Ctx.WriteString("恭喜用户信息注册成功")
}
