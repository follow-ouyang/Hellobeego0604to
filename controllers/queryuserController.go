package controllers

import (
	"HelloBeego0604to/db_mysql"
	"HelloBeego0604to/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type QueryuserController struct {
	beego.Controller
}

func (q *QueryuserController) Post() {


	//解析前端提交的数据
	//接收数据
	dataBytes,err := ioutil.ReadAll(q.Ctx.Request.Body)
	if err != nil {
		result := models.Result{
			COde:    0,
			Message: "数据接收错误，请重试",
			Data:    nil,
		}
		q.Data["json"] = &result
		q.ServeJSON()
		return
	}
	//fmt.Println(dataBytes)
	var name models.Names
	err = json.Unmarshal(dataBytes,&name)
	if err != nil {
		result := models.Result{
			COde:    0,
			Message: "数据解析失败，请重新输入后重试...",
			Data:    nil,
		}
		q.Data["json"] = &result
		q.ServeJSON()
		return
	}
	fmt.Println("用户名是：",name.User)
	q.Ctx.WriteString("用户名是："+name.User)

	query,err := db_mysql.QuerMoviesNum()
	if err != nil {
		fmt.Println()
	}

	//var message models.User
	//err = json.Unmarshal(dataBytes,&message)
	//if err != nil {
	//	result := models.Result{
	//		COde:    0,
	//		Message: "数据解析失败，请重试",
	//		Data:    nil,
	//	}
	//	q.Data["json"] = &result
	//	q.ServeJSON()
	//	return
	//}

	//row,err := db_mysql.AddUser(message)
	//if err != nil {
	//	result := models.Result{
	//		COde:    0,
	//		Message: "用户信息注册失败，请重试。。。",
	//		Data:    nil,
	//	}
	//	q.Data["json"] = &result
	//	q.ServeJSON()
	//	return
	//}
	//fmt.Println(row)
	//
	//md5Hash := md5.New()
	//md5Hash.Write([]byte(message.Password))
	//message.Password = hex.EncodeToString(md5Hash.Sum(nil))
	//result := models.Result{
	//	COde:    1,
	//	Message: "恭喜用户信息注册成功",
	//	Data:    message,
	//}
	//q.Data["json"] = &result
	//q.ServeJSON()

}
