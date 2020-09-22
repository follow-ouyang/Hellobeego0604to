package routers

import (
	"HelloBeego0604to/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
	//Router打一个参数为请求的路径，请求的路径必须保持唯一
    beego.Router("/index", &controllers.MainController{})
}
