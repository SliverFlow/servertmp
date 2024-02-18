package system

import (
	"github.com/gin-gonic/gin"
	"server/internal/service/system"
)

type UserApi struct {
	sysUserService *system.UserService
}

func NewSysUserApi(sysUserService *system.UserService) *UserApi {
	return &UserApi{
		sysUserService: sysUserService,
	}
}

// InitRouter 初始化 系统用户 相关路由
func (ua *UserApi) InitRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")

	{
		userApi.POST("/find", ua.sysUserService.Find)
		userApi.POST("/create", ua.sysUserService.Create)
		userApi.POST("/delete", ua.sysUserService.Delete)
		userApi.POST("/update", ua.sysUserService.Update)
	}
}
