package initialize

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ctx = context.Background()

func InitRedis() {
	rdbConfig := global.Config.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", rdbConfig.Host, rdbConfig.Port),
		Password: rdbConfig.Password, // No password set
		DB:       rdbConfig.Database, // Use default DB
		PoolSize: rdbConfig.PoolSize,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		global.Logger.Log(zapcore.ErrorLevel, "Ping to redis server failed!", zap.Error(err), zap.Bool("success", false))
	}

	fmt.Println("Init Redis is running")

	global.Rdb = client

	global.Logger.Info("Redis connected successfully!", zap.Bool("success", true))

	RedisExample()

}

func RedisExample() {
	err := global.Rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		global.Logger.Error("Redis | set key failed!", zap.Error(err), zap.Bool("success", false))
		panic(err)
	}

	value, err := global.Rdb.Get(ctx, "foo").Result()
	if err != nil {
		global.Logger.Error("Redis | get key failed!", zap.Error(err), zap.Bool("success", false))
		panic(err)
	}

	global.Logger.Info("Value of foo is::", zap.String("foo", value))
}
