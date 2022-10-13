package client_application

import (
	"errors"
	"fmt"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/UpdateClient"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Update(req *UpdateClient.Req) error {

	detail, err := Repository.ClientRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	var client Model.Client

	if req.Name != "" && req.Name != detail.Name {
		client.Name = req.Name
	}
	if req.Mark != "" && req.Mark != detail.Mark {
		client.Mark = req.Mark
	}
	if req.Remark != "" && req.Remark != detail.Remark {
		detailByMark, err := Repository.ClientRepository{}.FindByMark(req.Mark)
		if err != nil {
			return err
		}
		if detailByMark.ID != 0 {
			return errors.New(fmt.Sprintf("唯一标识符：%s，已存在。", req.Mark))
		}

		client.Remark = req.Remark
	}

	return Repository.ClientRepository{}.Update(&client, req.ID)
}
