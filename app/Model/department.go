package Model

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	gorm.Model
	UUID      string    `gorm:"type:char(32);not null;default:'';unique;comment:部门UUID" json:"uuid"`
	Name      string    `gorm:"type:varchar(64);not null;default:'';comment:部门名称" json:"name"`
	Mark      string    `gorm:"type:varchar(32);not null;default:'';unique;comment:部门唯一标识" json:"mark"`
	Remark    string    `gorm:"type:varchar(255);default:'';comment:备注" json:"remark"`
	EnableAt  time.Time `gorm:"type:datetime;default:null;comment:启用时间" format:"2006-01-02 15:04:05" json:"enable_at"`
	ForbadeAt time.Time `gorm:"type:datetime;default:null;comment:禁用时间" format:"2006-01-02 15:04:05" json:"forbade_at"`
}

func (Department) TableName() string {
	return "uc_department"
}
