package main

import (
	"go.uber.org/zap"
	"server/internal/config"
	"server/internal/core/initialize"
)

const path = "./etc/config.yaml"

func main() {
	// 初始化配置文件
	conf := config.Init(path)
	// 初始化日志
	log := initialize.Zap(conf)

	r := wireApp(conf, log)
	err := r.Server.ListenAndServe()
	if err != nil {
		log.Error("服务启动失败", zap.Error(err))
	}
}
