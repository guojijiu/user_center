package Model

import (
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	gorm.Model
	ClientID      uint      `gorm:"type:int(11);not null;default:0;index:idx_client_id;comment:所属客户端" json:"client_id"`
	Mark          string    `gorm:"type:varchar(32);not null;default:'';unique;comment:权限唯一标识" json:"mark"`
	Name          string    `gorm:"type:varchar(16);not null;default:'';comment:权限名称" json:"name"`
	Sort          uint      `gorm:"type:int(11);default:0;comment:排序序号" json:"sort"`
	Type          uint      `gorm:"type:tinyint(1);not null;default:0;comment:权限类型，1-目录，2-菜单，3-按钮，4-接口" json:"type"`
	ParentID      uint      `gorm:"type:int(11);not null;default:0;comment:父级id" json:"parent_id"`
	Remark        string    `gorm:"type:varchar(255);default:'';comment:备注" json:"remark"`
	IconPath      string    `gorm:"type:varchar(64);default:'';comment:图标路径" json:"icon_path"`
	RouteName     string    `gorm:"type:varchar(32);default:'';comment:路由名称" json:"route_name"`
	RoutePath     string    `gorm:"type:varchar(64);default:'';comment:路由路径" json:"route_path"`
	ModulePath    string    `gorm:"type:varchar(64);default:'';comment:组件路径" json:"module_path"`
	RequestMethod string    `gorm:"type:varchar(32);default:'';comment:请求方式" json:"request_method"`
	HiddenAt      time.Time `gorm:"type:datetime;default:null;comment:隐藏时间" json:"hidden_at"`
}

func (Permission) TableName() string {
	return "uc_permission"
}
