package controllers

import (
	"github.com/astaxie/beego"
)

type RabmqController struct {
	beego.Controller
}

func (c *RabmqController) TestRabMQ() {
	c.TplName = "index.html"
}
