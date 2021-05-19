package main

import (
	_ "golangwangjie/initial"
	_ "golangwangjie/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
