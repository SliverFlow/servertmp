package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/internal/biz"
	"server/internal/model/request"
	"server/pkg/response"
)

type UserService struct {
	log *zap.Logger

	systemUserUsercase *biz.SystemUserUsecase
}

func NewUserService(
	log *zap.Logger,
	systemUserUsercase *biz.SystemUserUsecase,
) *UserService {
	return &UserService{
		log:                log,
		systemUserUsercase: systemUserUsercase,
	}
}

func (us *UserService) Create(c *gin.Context) {
	var req request.UserCreateReq
	if err := c.ShouldBind(&req); err != nil {
		us.log.Error("[UserService] failed to bind request", zap.Error(err))
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	err := us.systemUserUsercase.Create(c, &req)
	if err != nil {
		us.log.Error("[UserService] failed to create user", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (us *UserService) Delete(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBind(&req); err != nil {
		us.log.Error("[UserService] failed to bind request", zap.Error(err))
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	err := us.systemUserUsercase.Delete(c, &req)
	if err != nil {
		us.log.Error("[UserService] failed to delete user", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (us *UserService) Update(c *gin.Context) {
	var req request.UserUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		us.log.Error("[UserService] failed to bind request", zap.Error(err))
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	err := us.systemUserUsercase.Update(c, &req)
	if err != nil {
		us.log.Error("[UserService] failed to update user", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (us *UserService) Find(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBind(&req); err != nil {
		us.log.Error("[UserService] failed to bind request", zap.Error(err))
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	user, err := us.systemUserUsercase.Find(c, &req)
	if err != nil {
		us.log.Error("[UserService] failed to find user", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}
