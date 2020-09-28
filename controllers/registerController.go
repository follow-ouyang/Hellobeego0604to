package controllers

import (
	"HelloBeego0604to/db_mysql"
	"HelloBeego0604to/models"
	"crypto/md5"
	"encoding/hex"
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
		//r.Ctx.WriteString("数据接收失败，请重试")
		result := models.Result{
			COde:    0,
			Message: "数据接收失败，请重试",
			Data:    nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
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
		//r.Ctx.WriteString("注册用户信息失败，请重试")
		result := models.Result{
			COde:    0,
			Message: "注册用户信息失败",
			Data:    nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
		return
	}
	fmt.Println(row)

	md5Hash := md5.New()
	md5Hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(md5Hash.Sum(nil))
	result := models.Result{
		COde:    1,
		Message: "恭喜用户信息注册成功",
		Data:    user,
	}
	//json.Marshal()   编码
	r.Data["json"] = &result
	r.ServeJSON()//将result编码为json格式返回给前端
	//r.Ctx.WriteString("恭喜用户信息注册成功")


}

