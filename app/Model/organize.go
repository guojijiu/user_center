package Model

import (
	"gorm.io/gorm"
	"time"
)

type Organize struct {
	gorm.Model
	UUID      string    `gorm:"type:char(32);not null;default:'';unique;comment:组织UUID" json:"uuid"`
	Name      string    `gorm:"type:varchar(64);not null;default:'';comment:组织名称" json:"name"`
	Mark      string    `gorm:"type:varchar(32);not null;default:'';unique;comment:组织唯一标识" json:"mark"`
	Remark    string    `gorm:"type:varchar(255);default:'';comment:备注" json:"remark"`
	EnableAt  time.Time `gorm:"type:datetime;default:null;comment:启用时间" format:"2006-01-02 15:04:05" json:"enable_at"`
	ForbadeAt time.Time `gorm:"type:datetime;default:null;comment:禁用时间" format:"2006-01-02 15:04:05" json:"forbade_at"`
}

func (Organize) TableName() string {
	return "uc_organize"
}
