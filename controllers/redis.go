package controllers

import (
	"context"
	"fmt"

	"golangwangjie/utils"

	"github.com/astaxie/beego"
)

type RedisController struct {
	beego.Controller
}

func (c *RedisController) TestRedis() {
	var ctx = context.Background()
	fmt.Println("test redis ...")
	utils.RedisSetCache(ctx, "key1", "value1", 10)
	val, _ := utils.RedisGetCache(ctx, "key1")
	fmt.Println("redis get key1 value ...", val)

	val2, _ := utils.RedisGetCache(ctx, "key2")
	fmt.Println("redis get key2 value ...", val2)
	c.TplName = "index.html"
}

func (c *RedisController) TestRedisDel() {
	var ctx = context.Background()
	fmt.Println("test redis del ...")
	utils.RedisDelCache(ctx, "key1")
	val2, _ := utils.RedisGetCache(ctx, "key1")
	fmt.Println("redis get key1 value ...", val2)
	c.TplName = "index.html"
}
