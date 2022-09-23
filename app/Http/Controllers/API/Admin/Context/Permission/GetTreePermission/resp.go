package GetTreePermission

import (
	"user_center/app/Model"
)

type Resp struct {
	ID       uint   `comment:"权限id" json:"id"`
	ClientID uint   `comment:"客户端id" json:"client_id"`
	Name     string `comment:"名称" json:"name"`
	ParentID uint   `comment:"父级id" json:"parent_id"`
}

func Item(permission []Model.Permission) []Resp {
	var list []Resp
	for _, v := range permission {
		var info Resp
		info.ID = v.ID
		info.ClientID = v.ClientID
		info.Name = v.Name
		info.ParentID = v.ParentID
		list = append(list, info)
	}
	return list
}
