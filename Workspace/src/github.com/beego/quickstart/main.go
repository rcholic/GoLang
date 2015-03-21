package main

import (
	"github.com/astaxie/beego"
	//_ "github.com/beego/quickstart/routers"
)

func main() {
	beego.SetStaticPath("/views", "views")
//    beego.Router("/users", &routers.UserController{})
	beego.Run()
}
