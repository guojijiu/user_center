package DetailRole

import "user_center/app/Model"

type Resp struct {
	Name   string `comment:"角色名称" json:"name"`
	Sort   uint   `comment:"排序序号" json:"sort"`
	Mark   string `comment:"角色唯一标识" json:"mark"`
	Remark string `comment:"备注" json:"remark"`
}

func Item(role *Model.Role) Resp {
	return Resp{
		Name:   role.Name,
		Sort:   role.Sort,
		Mark:   role.Mark,
		Remark: role.Remark,
	}
}
