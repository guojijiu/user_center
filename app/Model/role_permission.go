package Model

import (
	"time"
)

type RolePermission struct {
	ID           uint `gorm:"primary_key"`
	RoleID       uint `gorm:"type:int(11);not null;default:0;index:idx_role_permission;comment:角色id" json:"role_id"`
	PermissionID uint `gorm:"type:int(11);not null;default:0;index:idx_role_permission;comment:权限id" json:"permission_id"`
	CreatedAt    time.Time
}

func (RolePermission) TableName() string {
	return "uc_mapping_role_permission"
}
