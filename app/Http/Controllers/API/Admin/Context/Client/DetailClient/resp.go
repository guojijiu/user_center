package DetailClient

import "user_center/app/Model"

type Resp struct {
	Name   string `comment:"账号" json:"name"`
	Mark   string `comment:"手机号" json:"mark"`
	Remark string `comment:"邮箱" json:"remark"`
}

func Item(client *Model.Client) Resp {
	return Resp{
		Name:   client.Name,
		Mark:   client.Mark,
		Remark: client.Remark,
	}
}
