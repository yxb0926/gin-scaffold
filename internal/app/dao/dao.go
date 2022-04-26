package dao

import (
	"github.com/google/wire"
	"yxb.com/gin-scaffold/v2/internal/app/dao/user"
)

var Set = wire.NewSet(
	user.RepoSet,
)

type UserRepo = user.Repo
