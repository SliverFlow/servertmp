package api

import (
	"github.com/google/wire"
	"server/api/system"
)

var ProviderSet = wire.NewSet(
	// 注入总路由
	NewApiGroup,
	// 注入系统路由
	system.NewSystemApi,
	system.NewUserApi,

	// 注入用户端路由
	// xuye.NewXuyeApi,
)
