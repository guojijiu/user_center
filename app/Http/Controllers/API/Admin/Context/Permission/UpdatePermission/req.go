package UpdatePermission

type Req struct {
	ID            uint   `binding:"required" comment:"权限id" json:"id"`
	Name          string `binding:"max=16,min=1" comment:"权限名称" json:"name"`
	Sort          uint   `binding:"max=999,min=1" validate:"required,email" comment:"排序序号" json:"sort"`
	Mark          string `binding:"max=32,min=1" comment:"唯一标识" json:"mark"`
	Remark        string `validate:"max=255,min=1" comment:"备注" json:"remark"`
	Type          uint   `binding:"oneof=1 2 3 4" comment:"权限类型" json:"type"`
	ParentID      uint   `comment:"父级id" json:"parent_id"`
	IconPath      string `validate:"omitempty,min=1,max=64" comment:"图标路径" json:"icon_path"`
	RouteName     string `validate:"omitempty,min=1,max=32" comment:"路由名称" json:"route_name"`
	RoutePath     string `validate:"omitempty,min=1,max=64" comment:"路由路径" json:"route_path"`
	ModulePath    string `validate:"omitempty,min=1,max=64" comment:"组件路径" json:"module_path"`
	RequestMethod string `validate:"omitempty,min=1,max=32" comment:"请求方式" json:"request_method"`
	HiddenAt      string `comment:"隐藏时间" json:"hidden_at" time_format:"2006-01-02 15:04:05"`
}
