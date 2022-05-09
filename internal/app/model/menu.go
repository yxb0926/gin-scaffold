package model

import "fmt"

type Menu struct {
	BaseModel
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Router   string `json:"router"`
	ParentId int    `json:"parentId"`
	IsShow   int    `json:"isShow"`
	Sequence int    `json:"sequence"`
}

func (a *Menu) TableName() string {
	return "g_menu"
}

func (a *Menu) String() string {
	return fmt.Sprintf("Name: %s", a.Name)
}
