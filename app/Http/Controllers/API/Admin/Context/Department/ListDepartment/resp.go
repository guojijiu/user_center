package ListDepartment

import "user_center/app/Model"

type Resp struct {
	ID     uint   `comment:"部门id" json:"id"`
	Name   string `comment:"名称" json:"name"`
	Mark   string `comment:"唯一标识符" json:"mark"`
	Remark string `comment:"备注" json:"remark"`
}

func Item(model []Model.Department) []Resp {
	var list []Resp
	for _, v := range model {
		var info Resp
		info.ID = v.ID
		info.Name = v.Name
		info.Mark = v.Mark
		info.Remark = v.Remark
		list = append(list, info)
	}
	return list
}
