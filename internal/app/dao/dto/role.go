package dto

import "yxb.com/gin-scaffold/v2/internal/app/model"

type RolePageReq PageReq

type RolePageRes struct {
	List *[]model.Role
	PageRes
}
