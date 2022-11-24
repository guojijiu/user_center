package DetailByForget

import "user_center/app/Model"

type Resp struct {
	ID      uint   `comment:"主键ID" json:"name"`
	Account string `comment:"账号" json:"mark"`
	Email   string `comment:"邮箱" json:"remark"`
	Phone   string `comment:"手机" json:"phone"`
}

func Item(model *Model.UserAuth) Resp {
	return Resp{
		ID:      model.ID,
		Account: model.Account,
		Email:   model.Email,
		Phone:   model.Phone,
	}
}
