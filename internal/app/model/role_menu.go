package model

type RoleMenu struct {
	BaseModel
	RoleId int `json:"roleId"`
	MenuId int `json:"menuId"`
}

func (a *RoleMenu) TableName() string {
	return "g_role_menu"
}
