package DetailPermission

import (
	"time"
	"user_center/app/Model"
)

type Resp struct {
	Name          string    `comment:"名称" json:"name"`
	Sort          uint      `comment:"排序序号" json:"sort"`
	Mark          string    `comment:"唯一标识" json:"mark"`
	Remark        string    `comment:"备注" json:"remark"`
	Type          uint      `comment:"类型" json:"type"`
	ParentID      uint      `comment:"父级id" json:"parent_id"`
	IconPath      string    `comment:"图标路径" json:"icon_path"`
	RouteName     string    `comment:"路由名称" json:"route_name"`
	RoutePath     string    `comment:"路由路径" json:"route_path"`
	ModulePath    string    `comment:"组件路径" json:"module_path"`
	RequestMethod string    `comment:"请求方式" json:"request_method"`
	HiddenAt      time.Time `comment:"隐藏时间" json:"hidden_at"`
}

func Item(permission *Model.Permission) Resp {
	return Resp{
		Name:          permission.Name,
		Sort:          permission.Sort,
		Mark:          permission.Mark,
		Remark:        permission.Remark,
		Type:          permission.Type,
		ParentID:      permission.ParentID,
		IconPath:      permission.IconPath,
		RouteName:     permission.RouteName,
		RoutePath:     permission.RoutePath,
		ModulePath:    permission.ModulePath,
		RequestMethod: permission.RequestMethod,
		HiddenAt:      permission.HiddenAt,
	}
}
