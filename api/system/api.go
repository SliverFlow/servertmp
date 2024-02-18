package system

import "github.com/gin-gonic/gin"

type ApiGroup struct {
	userApi *UserApi
}

func NewSystemApi(userApi *UserApi) *ApiGroup {
	return &ApiGroup{
		userApi: userApi,
	}
}

// InitRouter 初始化系统相关路由
func (sa *ApiGroup) InitRouter(router *gin.Engine) {

	systemApiRouter := router.Group("/system")

	sa.userApi.InitRouter(systemApiRouter)
}
