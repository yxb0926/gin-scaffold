package dto

import "yxb.com/gin-scaffold/v2/internal/app/model"

type UserPageReq PageReq

type UserPageRes struct {
	List *[]model.User `json:"userList"`
	PageRes
}
