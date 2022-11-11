package DetailOrganize

import "user_center/app/Model"

type Resp struct {
	Name   string `comment:"名称" json:"name"`
	Mark   string `comment:"标识符" json:"mark"`
	Remark string `comment:"备注" json:"remark"`
}

func Item(model *Model.Organize) Resp {
	return Resp{
		Name:   model.Name,
		Mark:   model.Mark,
		Remark: model.Remark,
	}
}
