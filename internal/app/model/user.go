package model

import (
	"fmt"
)

type User struct {
	BaseModel
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	Creator  string `json:"creator"`
}

func (a *User) TableName() string {
	return "g_user"
}

func (a *User) String() string {
	return fmt.Sprintf("id:%d, userName:%s, email:%s, status:%d, creator:%s, createdAt:%s, updatedAt:%s",
		a.ID, a.UserName, a.Email, a.Status, a.Creator, a.CreatedAt, a.UpdatedAt)
}
