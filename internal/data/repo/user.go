package repo

import (
	"context"
	"server/internal/model"
)

type IUserRepo interface {
	Create(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user *model.User) error
	Find(ctx context.Context, id int64) (*model.User, error)
}
