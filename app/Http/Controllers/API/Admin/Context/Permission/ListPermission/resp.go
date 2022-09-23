package ListPermission

import (
	"time"
	"user_center/app/Model"
)

type Resp struct {
	ID        uint      `comment:"用户id" json:"id"`
	ClientID  uint      `comment:"所属客户端" json:"client_id"`
	Name      string    `comment:"角色名称" json:"name"`
	Sort      uint      `comment:"排序序号" json:"sort"`
	Mark      string    `comment:"角色唯一标识" json:"mark"`
	Remark    string    `comment:"备注" json:"remark"`
	Type      uint      `comment:"类型" json:"type"`
	ParentID  uint      `comment:"父级id" json:"parent_id"`
	HiddenAt  time.Time `comment:"隐藏时间" json:"hidden_at"`
	CreatedAt time.Time `comment:"创建时间" json:"created_at"`
	UpdatedAt time.Time `comment:"更新时间" json:"updated_at"`
}

func Item(permission []Model.Permission) []Resp {
	var list []Resp
	for _, v := range permission {
		var info Resp
		info.ID = v.ID
		info.ClientID = v.ClientID
		info.Name = v.Name
		info.Sort = v.Sort
		info.Mark = v.Mark
		info.Remark = v.Remark
		info.Type = v.Type
		info.ParentID = v.ParentID
		info.HiddenAt = v.HiddenAt
		info.CreatedAt = v.CreatedAt
		info.UpdatedAt = v.UpdatedAt
		list = append(list, info)
	}
	return list
}
