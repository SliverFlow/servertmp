package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"server/internal/config"
	"server/internal/core/initialize"
)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData(config *config.Possess) *Data {

	db := initialize.InitMysqlConn(config)
	rdb := initialize.InitRedisConn(config)

	return &Data{
		db:  db,
		rdb: rdb,
	}
}

func (c *Data) DB(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}

func (c *Data) RDB() *redis.Client {
	return c.rdb
}
