package data

import (
	"github.com/google/wire"
	"server/internal/data/repo"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	wire.Bind(new(repo.IUserRepo), new(*userRepo)),
)
