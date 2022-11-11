package Model

import (
	"time"
)

type ClientOrganize struct {
	ID         uint `gorm:"primary_key"`
	ClientID   uint `gorm:"type:int(11);not null;default:0;index:idx_client_organize;comment:客户端id" json:"client_id"`
	OrganizeID uint `gorm:"type:int(11);not null;default:0;index:idx_client_organize;comment:组织id" json:"organize_id"`
	CreatedAt  time.Time
}

func (ClientOrganize) TableName() string {
	return "uc_mapping_client_organize"
}
