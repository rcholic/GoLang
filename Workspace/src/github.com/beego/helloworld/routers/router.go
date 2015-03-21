package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/helloworld/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{})
}
