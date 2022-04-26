package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserPagReq PageReq

type UserPageRes struct {
	List *[]User `json:"userList"`
	PageRes
}

func (a *User) TableName() string {
	return "g_user"
}

func (a *User) ToString() string {
	return fmt.Sprintf("id:%d, userName:%s, email:%s, status:%d, creator:%s, createdAt:%s, updatedAt:%s",
		a.ID, a.UserName, a.Email, a.Status, a.Creator, a.CreatedAt, a.UpdatedAt)
}
