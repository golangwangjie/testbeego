package controllers

import (
	"fmt"

	"golangwangjie/initial"
	"golangwangjie/utils"

	"github.com/astaxie/beego"
)

type ZKController struct {
	beego.Controller
}

func (c *ZKController) TestZookeepers() {
	fmt.Println("test zookeepers ...")
	utils.ZKadd("/test", []byte("zk-test-value"))
	utils.ZKget("/test")

	utils.ZKmodify("/test", []byte("zk-new"))
	utils.ZKget("/test")

	utils.ZKdel("/test")
	utils.ZKget("/test")

	//测试锁，用的父节点
	utils.ZKadd("/lock", []byte("zklock"))
	c.TplName = "index.html"
}

func (c *ZKController) TestZookeepersLock() {
	utils.ZKLock("/lock/lock001")
	c.TplName = "index.html"
}

func (c *ZKController) TestZookeepersSub() {
	_, _, event, err := initial.ZKConn.ExistsW("/watchzk")
	if err != nil {
		fmt.Println(err)
		return
	}
	//订阅 节点 /watchzk
	utils.ZKsubscribe(event)
	fmt.Println("sub end ...")
	c.TplName = "index.html"
}

func (c *ZKController) TestZookeepersPub() {
	//给节点 /watchzk 发布消息
	utils.ZKpublish("/watchzk", []byte("watch 001"))
	fmt.Println("pub end ...")
	c.TplName = "index.html"
}
