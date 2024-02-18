package initialize

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"server/internal/config"
	"time"
)

func InitMysqlConn(c *config.Possess) *gorm.DB {
	conf := c.Mysql
	dsn := conf.Dsn()

	mc := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	logLever := conf.GetLog()

	DB, err := gorm.Open(mysql.New(mc), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second,               // 慢 SQL 阈值
				LogLevel:      logger.LogLevel(logLever), // Log level
				Colorful:      false,                     // 禁用彩色打印
			}),
	})

	if err != nil {
		panic(err)
	}

	db, _ := DB.DB()
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	// DB.AutoMigrate(&model.User{})

	return DB
}
