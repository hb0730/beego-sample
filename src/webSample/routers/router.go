package routers

import (
	"github.com/astaxie/beego"
	"webSample/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.TestController{}, "get:Get")
}
