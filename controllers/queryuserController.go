package controllers

import (
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

	var monicker models.Monicker
	err = json.Unmarshal(dataBytes,&monicker)
	if err != nil {
		result := models.Result{
			COde:    0,
			Message: "数据解析失败，请重试",
			Data:    nil,
		}
		q.Data["json"] = &result
		q.ServeJSON()
		fmt.Println(err.Error())
		return
	}

}
