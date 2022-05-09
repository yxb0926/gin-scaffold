package model

import "fmt"

type UserRole struct {
	BaseModel
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}

func (a *UserRole) TableName() string {
	return "g_user_role"
}

func (a *UserRole) String() string {
	return fmt.Sprintf("ID: %d, UserID: %d, RoleId: %d, CreatedAt:%s, UpdatedAt:%s",
		a.ID, a.UserId, a.RoleId, a.CreatedAt, a.UpdatedAt)
}
