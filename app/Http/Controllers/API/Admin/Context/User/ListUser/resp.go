package ListUser

import "user_center/app/Model"

type Resp struct {
	ID      uint   `comment:"用户id" json:"id"`
	Account string `comment:"账号" json:"account"`
	Phone   string `comment:"手机号" json:"phone"`
	Email   string `comment:"邮箱" json:"email"`
}

func Item(user []Model.UserAuth) []Resp {
	var list []Resp
	for _, v := range user {
		var info Resp
		info.ID = v.ID
		info.Account = v.Account
		info.Phone = v.Phone
		info.Email = v.Email
		list = append(list, info)
	}
	return list
}
