package initial

import (
	"context"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var ctx = context.Background()
var RedisClusterClient *redis.ClusterClient

func InitCache() {
	cacheConfig := beego.AppConfig.String("cache")
	if "redis" == cacheConfig {
		initRedis()
	} else {
		fmt.Println("Not Supported by now ...")
		// initMemcache()
	}
}

func initRedis() {
	fmt.Println("init redis cache ...")
	options := redis.ClusterOptions{
		Addrs: []string{"192.168.56.1:7001", "192.168.56.1:7002", "192.168.56.1:7003"},
	}
	// 新建一个client
	RedisClusterClient = redis.NewClusterClient(&options)
	RedisClusterClient.Ping(ctx)
	// close
	// defer RedisClusterClient.Close()
}
