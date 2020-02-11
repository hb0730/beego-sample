package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Admin struct {
	beego.Controller
}
type User struct {
	Id   string `form:"id"`
	Name string
}

func (c *Admin) Get() {
	c.Ctx.WriteString("hhhh")
}

// get 127.0.0.1:8080/admin/get?id=1&name=123
func (c *Admin) GetParams() {
	name := c.GetString("name")
	id := c.GetString("id")
	fmt.Printf("id = %s ,name = %s", id, name)
	c.Ctx.WriteString("id = " + id + " name = " + name)
}

// post 127.0.0.1:8080/admin/postParams
func (c *Admin) PostParams() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		panic(err)
	}
	c.Ctx.WriteString("user id = " + u.Id + " user name =" + u.Name)
}

// post 127.0.0.1:8080/admin/getHtml
func (c *Admin) GetHtml() {
	c.Ctx.WriteString("<html>" +
		"<form id='user' action='http://127.0.0.1:8080/admin/postParams' method='post'>" +
		"id: <input name='id' type='text'>" +
		"name: <input name='Name' type='text'>" +
		"<input type='submit' value='提交'>" +
		" </form>" +
		"</html")
}
