package routers

import (
	"github.com/astaxie/beego"
	"webSample/controllers"
)

func init() {
	beego.Router("/admin", &controllers.Admin{})
	beego.Router("/admin/get", &controllers.Admin{}, "get:GetParams")
	beego.Router("/admin/getHtml", &controllers.Admin{}, "get:GetHtml")
	beego.Router("/admin/postParams", &controllers.Admin{}, "post:PostParams")
}
