package utils

import (
	"context"
	"fmt"
	"golangwangjie/initial"
	"time"

	"github.com/go-redis/redis"
)

//string类型
//缓存。。。
//设置
func RedisSetCache(ctx context.Context, key string, value interface{}, timeout int) error {
	timeouts := time.Duration(timeout) * time.Second
	err := initial.RedisClusterClient.Set(ctx, key, value, timeouts).Err()
	return checkErr(err, "redis set value")
}

//获取
func RedisGetCache(ctx context.Context, key string) (string, error) {
	val, err := initial.RedisClusterClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("redis get value error ...", err)
		return "error", err
	}
	return val, nil
}

//删除
func RedisDelCache(ctx context.Context, key string) error {
	err := initial.RedisClusterClient.Del(ctx, key).Err()
	return checkErr(err, "redis del value")
}

//bitmap
//设置偏移量的值
func RedisSetBitmapCache(ctx context.Context, key string, offset int64, value int) error {
	err := initial.RedisClusterClient.SetBit(ctx, key, offset, value).Err()
	return checkErr(err, "redis bitmap")
}

//某段时间的1，（例如用户签到）
func RedisGetBitmapCount(ctx context.Context, key string, startend *redis.BitCount) error {
	err := initial.RedisClusterClient.BitCount(ctx, key, startend).Err()
	return checkErr(err, "redis bitmap count")
}

//统计活跃用户
