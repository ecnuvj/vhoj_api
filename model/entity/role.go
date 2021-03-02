package entity

type Role struct {
	RoleId   uint   `json:"role_id" form:"role_id" binding:"required"`
	RoleName string `json:"role_name" form:"role_name" binding:"required"`
}
