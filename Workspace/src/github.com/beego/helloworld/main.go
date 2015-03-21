package main

import (
	_ "github.com/beego/helloworld/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}

