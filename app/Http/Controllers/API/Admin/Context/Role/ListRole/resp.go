package ListRole

import "user_center/app/Model"

type Resp struct {
	ID       uint   `comment:"用户id" json:"id"`
	ClientID uint   `comment:"所属客户端" json:"client_id"`
	Name     string `comment:"角色名称" json:"name"`
	Sort     uint   `comment:"排序序号" json:"sort"`
	Mark     string `comment:"角色唯一标识" json:"mark"`
	Remark   string `comment:"备注" json:"remark"`
}

func Item(role []Model.Role) []Resp {
	var list []Resp
	for _, v := range role {
		var info Resp
		info.ID = v.ID
		info.ClientID = v.ClientID
		info.Name = v.Name
		info.Sort = v.Sort
		info.Mark = v.Mark
		info.Remark = v.Remark
		list = append(list, info)
	}
	return list
}
