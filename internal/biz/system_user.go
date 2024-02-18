package biz

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/internal/data/repo"
	"server/internal/model"
	"server/internal/model/request"
	"server/pkg/zerror"
	"time"
)

type SystemUserUsecase struct {
	log *zap.Logger

	userRepo repo.IUserRepo
}

func NewSysUserUsecase(
	log *zap.Logger,
	userRepo repo.IUserRepo,
) *SystemUserUsecase {
	return &SystemUserUsecase{
		log:      log,
		userRepo: userRepo,
	}
}

func (uu *SystemUserUsecase) Create(ctx context.Context, req *request.UserCreateReq) (err error) {
	userCreate := &model.User{
		WxOpenId:   "r23r23",
		Username:   req.Username,
		Password:   req.Password,
		Status:     model.UserStatusNormal,
		DeleteFlag: model.ModelNotDeleted,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	err = uu.userRepo.Create(ctx, userCreate)
	if err != nil {
		uu.log.Error("[SystemUserUsecase] failed to create user", zap.Error(err))
		return zerror.NewWithMessage("用户创建失败")
	}
	return
}

func (uu *SystemUserUsecase) Delete(ctx context.Context, req *request.IdReq) (err error) {
	err = uu.userRepo.Delete(ctx, req.Id)
	if err != nil {
		uu.log.Error("[SystemUserUsecase] failed to delete user", zap.Error(err))
		return zerror.NewWithMessage("用户删除失败")
	}
	return
}

func (uu *SystemUserUsecase) Update(ctx context.Context, req *request.UserUpdateReq) (err error) {
	userUpdate := &model.User{
		Id:         req.Id,
		Username:   req.Username,
		Password:   req.Password,
		UpdateTime: time.Now().Unix(),
	}
	err = uu.userRepo.Update(ctx, userUpdate)
	if err != nil {
		uu.log.Error("[SystemUserUsecase] failed to update user", zap.Error(err))
		return zerror.NewWithMessage("用户更新失败")
	}
	return
}

func (uu *SystemUserUsecase) Find(ctx context.Context, req *request.IdReq) (user *model.User, err error) {
	user, err = uu.userRepo.Find(ctx, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, zerror.NewWithMessage("用户未找到")
		}
		uu.log.Error("[SystemUserUsecase] failed to find user", zap.Error(err))
		return nil, zerror.NewWithMessage("用户查找失败")
	}
	return user, nil
}
