package dto

import "yxb.com/gin-scaffold/v2/internal/app/model"

type MenuPageReq PageReq

type MenuPageRes struct {
	List *[]model.Menu
	PageRes
}
