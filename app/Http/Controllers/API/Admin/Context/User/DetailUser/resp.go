package DetailUser

import "user_center/app/Model"

type Resp struct {
	Account string `comment:"账号" json:"account"`
	Phone   string `comment:"手机号" json:"phone"`
	Email   string `comment:"邮箱" json:"email"`
}

func Item(user *Model.UserAuth) Resp {
	return Resp{
		Account: user.Account,
		Phone:   user.Phone,
		Email:   user.Email,
	}
}
