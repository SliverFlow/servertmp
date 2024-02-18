package xuye

import "github.com/gin-gonic/gin"

type ApiGroup struct {
}

func NewXuyeApi() *ApiGroup {
	return &ApiGroup{}
}

// InitRouter 初始化系统相关路由
func (sa *ApiGroup) InitRouter(router *gin.Engine) {

	// xuyeApiRouter := router.Group("/xuye")

}
