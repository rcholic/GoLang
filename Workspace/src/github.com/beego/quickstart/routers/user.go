package routers

import (
	"github.com/astaxie/beego"
)

type UserContrller struct {
	beego.Controller
}

func (u *UserContrller) Get() {
	u.Data["username"] = "ivytony"
	u.Data["password"] = "iloveu"
	u.TplNames = "user.html"
}
