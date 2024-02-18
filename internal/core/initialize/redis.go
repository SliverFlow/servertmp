package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/internal/config"
)

func InitRedisConn(c *config.Possess) *redis.Client {
	conf := c.Redis
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})

	_, err := cli.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return cli
}
