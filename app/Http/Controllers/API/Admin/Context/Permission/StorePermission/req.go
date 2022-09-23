package StorePermission

type Req struct {
	ClientID      uint   `binding:"required" comment:"所属客户端" json:"client_id"`
	Name          string `binding:"required,max=16,min=1" comment:"权限名称" json:"name"`
	Sort          uint   `binding:"required,max=999,min=1" validate:"required,email" comment:"排序序号" json:"sort"`
	Mark          string `binding:"required,max=32,min=1" comment:"权限唯一标识" json:"mark"`
	Type          uint   `binding:"required,oneof=1 2 3 4" comment:"权限类型" json:"type"`
	ParentID      uint   `comment:"父级id" json:"parent_id"`
	Remark        string `validate:"omitempty,max=255,min=1" comment:"备注" json:"remark"`
	IconPath      string `validate:"omitempty,min=1,max=64" comment:"图标路径" json:"icon_path"`
	RouteName     string `validate:"omitempty,min=1,max=32" comment:"路由名称" json:"route_name"`
	RoutePath     string `validate:"omitempty,min=1,max=64" comment:"路由路径" json:"route_path"`
	ModulePath    string `validate:"omitempty,min=1,max=64" comment:"组件路径" json:"module_path"`
	RequestMethod string `validate:"omitempty,min=1,max=32" comment:"请求方式" json:"request_method"`
	HiddenAt      string `comment:"隐藏时间" json:"hidden_at" time_format:"2006-01-02 15:04:05"`
}
