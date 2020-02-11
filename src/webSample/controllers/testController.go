package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (test *TestController) Get() {
	test.Ctx.WriteString("welCome test controller")
}
