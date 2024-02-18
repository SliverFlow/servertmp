package service

import (
	"github.com/google/wire"
	"server/internal/service/system"
)

var ProviderSet = wire.NewSet(
	// 系统用户服务
	system.NewUserService,

	// 用户端用户服务
	// xuye.NewXuyeUserService,
)
