package model

import "fmt"

type Role struct {
	BaseModel
	Name   string `json:"name"`
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}

func (a *Role) TableName() string {
	return "g_role"
}

func (a *Role) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Status: %d, Desc: %s, CreatedAt: %s, UpdatedAt: %s",
		a.ID, a.Name, a.Status, a.Desc, a.CreatedAt, a.UpdatedAt)
}
