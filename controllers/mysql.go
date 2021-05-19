package controllers

import (
	"fmt"

	"golangwangjie/models"

	"github.com/astaxie/beego"
)

type MysqlController struct {
	beego.Controller
}

func (c *MysqlController) Get() {
	c.TplName = "index.html"
}

func (c *MysqlController) TestMysql() {
	fmt.Println("test mysql ...")
	result := models.GetActorInfo(1)
	fmt.Println("mysql result... = ", result)
	c.TplName = "index.html"
}

func (c *MysqlController) TestMysqlIndexUsers() {
	fmt.Println("Test Mysql Index Users ...")
	models.IndexUsers()
	models.AddUserInfo()
	result := models.GetUserInfo(1)
	fmt.Println("mysql users get first data ... = ", result)
	models.UpdateUserInfo("wangjie", 20)
	models.DelUserInfo("wangjie")
	c.TplName = "index.html"
}
