package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["username"] = "beego.me"
	u.Data["password"] = "astaxie@gmail.com"
	u.TplNames = "newuser.html"
}

func (this *UserController) Post() {
	//var result string
	username := this.Input().Get("username")
	this.Data["json"] = username
	this.ServeJson()
}
