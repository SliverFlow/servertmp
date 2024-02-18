package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/internal/config"
	"server/internal/core/initialize/internal"
	utils "server/pkg/util"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap(c *config.Possess) (logger *zap.Logger) {

	conf := c.Zap

	if ok, _ := utils.PathExists(conf.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", conf.Director)
		_ = os.Mkdir(conf.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores(conf)
	logger = zap.New(zapcore.NewTee(cores...))

	if conf.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
