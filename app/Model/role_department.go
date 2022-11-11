package Model

import (
	"time"
)

type RoleDepartment struct {
	ID           uint `gorm:"primary_key"`
	RoleID       uint `gorm:"type:int(11);not null;default:0;index:idx_role_department;comment:角色id" json:"role_id"`
	DepartmentID uint `gorm:"type:int(11);not null;default:0;index:idx_role_department;comment:部门id" json:"department_id"`
	CreatedAt    time.Time
}

func (RoleDepartment) TableName() string {
	return "uc_mapping_role_department"
}
