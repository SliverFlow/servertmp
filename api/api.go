package api

import (
	"github.com/gin-gonic/gin"
	"server/api/system"
)

type Group struct {
	systemApi *system.ApiGroup

	// xuyeApi *xuye.ApiGroup
}

func NewApiGroup(systemApi *system.ApiGroup) *Group {
	return &Group{
		systemApi: systemApi,
	}
}

// InitApi 初始化总路由
func (api *Group) InitApi(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 初始化系统路由
	api.systemApi.InitRouter(router)

	// 初始化用户端路由 (xuye)
	// api.xuyeApi.InitRouter(router)
}
