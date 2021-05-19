package routers

import (
	"golangwangjie/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MysqlController{})
	beego.Router("/testmysql", &controllers.MysqlController{}, "*:TestMysql")
	beego.Router("/testmysqlIndexusers", &controllers.MysqlController{}, "*:TestMysqlIndexUsers")

	beego.Router("/testredis", &controllers.RedisController{}, "*:TestRedis")
	beego.Router("/testredisdel", &controllers.RedisController{}, "*:TestRedisDel")

	beego.Router("/testzookeepers", &controllers.ZKController{}, "*:TestZookeepers")
	beego.Router("/testzookeeperslock", &controllers.ZKController{}, "*:TestZookeepersLock")
	beego.Router("/testzookeeperssub", &controllers.ZKController{}, "*:TestZookeepersSub")
	beego.Router("/testzookeeperspub", &controllers.ZKController{}, "*:TestZookeepersPub")
}
