package Model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UUID      string `gorm:"type:char(32);not null;default:'';unique;comment:'用户UUID'" json:"uuid"`
	Account   string `gorm:"type:varchar(32);not null;default:'';unique;comment:'登录账号'" json:"account"`
	Phone     string `gorm:"type:char(11);unique;comment:'手机号'" json:"phone"`
	Email     string `gorm:"type:varchar(32);unique;;comment:'邮箱'" json:"email"`
	Passwd    string `gorm:"type:varchar(128);not null;default:'';comment:'密码密文'" json:"-"`
	ForbadeAt *time.Time
}

func (User) TableName() string {
	return "uc_user_auth"
}
