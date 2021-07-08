package Model

import (
	"time"
)

type UserRole struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `gorm:"type:int(11);not null;default:0;index:idx_user_role;comment:'用户id'" json:"user_id"`
	RoleID    uint `gorm:"type:int(11);not null;default:0;index:idx_user_role;comment:'角色id'" json:"role_id"`
	CreatedAt time.Time
}

func (UserRole) TableName() string {
	return "uc_mapping_user_role"
}
