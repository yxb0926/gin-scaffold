package dao

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	UserRepoSet,
	RoleRepoSet,
	MenuRepoSet,
)
