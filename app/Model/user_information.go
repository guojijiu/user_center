package Model

import (
	"gorm.io/gorm"
	"time"
)

type UserInformation struct {
	gorm.Model
	UserID          uint       `gorm:"type:int(11);not null;default:0;unique;comment:用户ID" json:"user_id"`
	Nickname        string     `gorm:"type:varchar(64);default:'';comment:用户昵称" json:"nickname"`
	HeaderImgPath   string     `gorm:"type:varchar(255);default:'';comment:用户头像地址" json:"header_img_path"`
	Sex             string     `gorm:"type:char(1);default:'H';comment:性别，M：男；F：女；H：隐藏" json:"sex"`
	Birthday        *time.Time `gorm:"type:date;default:null;comment:出生年月日" format:"2006-01-02 15:04:05" json:"birthday"`
	Address         string     `gorm:"type:varchar(64);default:'';comment:地址" json:"address"`
	Organization    string     `gorm:"type:varchar(32);default:'';comment:公司，组织或院校" json:"organization"`
	PersonalProfile string     `gorm:"type:varchar(255);default:'';comment:个人简介" json:"personal_profile"`
}

func (UserInformation) TableName() string {
	return "uc_user_information"
}
