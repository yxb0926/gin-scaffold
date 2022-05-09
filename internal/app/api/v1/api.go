package v1

import "github.com/google/wire"

var ApiSet = wire.NewSet(
	UserApiSet,
	RoleApiSet,
	MenuApiSet,
)
