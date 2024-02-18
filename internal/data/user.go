package data

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"server/internal/config"
	"server/internal/model"
	"server/pkg/util"
	"server/pkg/zerror"
	"strconv"
	"time"
)

type userRepo struct {
	log *zap.Logger

	config        *config.Possess
	data          *Data
	rdbExpireTime time.Duration
}

func NewUserRepo(
	log *zap.Logger,
	config *config.Possess,
	data *Data,
) *userRepo {
	return &userRepo{
		log:           log,
		config:        config,
		data:          data,
		rdbExpireTime: time.Duration(config.Redis.ExpireTime) * time.Minute,
	}
}

// Create 创建用户
func (ur *userRepo) Create(ctx context.Context, user *model.User) (err error) {
	tx := ur.data.DB(ctx).Model(&model.User{})
	err = tx.Create(user).Error
	if err != nil {
		return errors.WithStack(err)
	}
	_ = ur.cacheSet(ctx, user)
	return nil
}

// Delete 删除用户
func (ur *userRepo) Delete(ctx context.Context, id int64) (err error) {
	tx := ur.data.DB(ctx).Model(&model.User{})

	um := make(map[string]interface{})
	um[model.UserCol.DeleteFlag] = model.ModelDeleted
	um[model.UserCol.UpdateTime] = time.Now().Unix()
	um[model.UserCol.DeleteTime] = time.Now().Unix()

	err = tx.Where(model.UserCol.DeleteFlag+" = ?", model.ModelDeleted).Where(model.UserCol.Status+" = ?", model.UserStatusNormal).Where(model.UserCol.Id+" = ?", id).Updates(um).Error
	if err != nil {
		return errors.WithStack(err)
	}
	_ = ur.cacheDel(ctx, id)
	return nil
}

// Update 更新用户
func (ur *userRepo) Update(ctx context.Context, user *model.User) (err error) {
	tx := ur.data.DB(ctx).Model(&model.User{})
	err = tx.Where(model.UserCol.DeleteFlag+" = ?", model.ModelNotDeleted).Where(model.UserCol.Status+" = ?", model.UserStatusNormal).Where(model.UserCol.Id+" = ?", user.Id).Updates(user).Error
	if err != nil {
		return errors.WithStack(err)
	}
	_ = ur.cacheDel(ctx, user.Id)
	return nil
}

// Find 查询用户
func (ur *userRepo) Find(ctx context.Context, id int64) (user *model.User, err error) {
	userCache, err := ur.CacheGet(ctx, id)
	if err == nil && userCache != nil {
		return userCache, nil
	}
	tx := ur.data.DB(ctx).Model(&model.User{})
	err = tx.Where(model.UserCol.DeleteFlag+" = ?", model.ModelNotDeleted).Where(model.UserCol.Status+" = ?", model.UserStatusNormal).Where(model.UserCol.Id+" = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	_ = ur.cacheSet(ctx, user)
	return user, nil
}

// getCacheKey 获取用户信息缓存key
func (ur *userRepo) getCacheKey(id int64) string {
	return ur.config.Redis.BaseKey + ":user:id:" + strconv.FormatInt(id, 10)
}

// cacheSet 设置用户信息缓存
func (ur *userRepo) cacheSet(ctx context.Context, user *model.User) error {
	if user == nil || user.Id == 0 {
		return zerror.NewWithMessage("user id is empty")
	}
	userStr, err := util.StructToJSON(user)
	if err != nil {
		ur.log.Error("[userRepo.CacheSet] util.StructToJSON", zap.Error(err))
		return errors.WithStack(err)
	}
	key := ur.getCacheKey(user.Id)
	return ur.data.rdb.Set(ctx, key, userStr, ur.rdbExpireTime).Err()
}

// cacheDel 删除用户信息缓存
func (ur *userRepo) cacheDel(ctx context.Context, id int64) error {
	key := ur.getCacheKey(id)
	return ur.data.rdb.Del(ctx, key).Err()
}

// CacheGet 获取用户信息缓存
func (ur *userRepo) CacheGet(ctx context.Context, id int64) (user *model.User, err error) {
	key := ur.getCacheKey(id)
	userStr, err := ur.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	err = util.JSONToStruct(userStr, &user)
	if err != nil {
		ur.log.Error("[userRepo.CacheGet] util.JSONToStruct", zap.Error(err))
		return nil, errors.WithStack(err)
	}
	ur.log.Info("[userRepo.CacheGet] cache 命中", zap.Any("user", user))
	return user, nil
}
