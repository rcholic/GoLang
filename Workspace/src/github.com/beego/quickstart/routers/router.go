package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/quickstart/controllers"
	"github.com/beego/quickstart/routers"
)

func init() {
	//beego.Include(&controllers.UserController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &routers.UserController{})
}
