package Model

import (
	"time"
)

type UserClient struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `gorm:"type:int(11);not null;default:0;index:idx_user_client;comment:用户id" json:"user_id"`
	ClientID  uint `gorm:"type:int(11);not null;default:0;index:idx_user_client;comment:客户端id" json:"client_id"`
	CreatedAt time.Time
}

func (UserClient) TableName() string {
	return "uc_mapping_user_client"
}
