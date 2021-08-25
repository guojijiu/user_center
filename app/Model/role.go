package Model

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	ClientID uint   `gorm:"type:int(11);not null;default:0;index:idx_client_id;comment:所属客户端" json:"client_id"`
	Name     string `gorm:"type:varchar(16);not null;default:'';comment:角色名称" json:"name"`
	Sort     uint   `gorm:"type:int(11);default:0;comment:排序序号" json:"sort"`
	Mark     string `gorm:"type:varchar(32);not null;default:'';unique;comment:角色唯一标识" json:"mark"`
	Remark   string `gorm:"type:varchar(255);default:'';comment:备注" json:"remark"`
}

func (Role) TableName() string {
	return "uc_role"
}
