package model

import "time"

type BaseModel struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PageReq struct {
	PageNum  int  `json:"pageNum"`
	PageSize int  `json:"pageSize"`
	Desc     bool `json:"desc`
}

type PageRes struct {
	Total    int64 `json:"total`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}
