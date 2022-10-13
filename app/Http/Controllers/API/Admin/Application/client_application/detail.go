package client_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/DetailClient"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Detail(req *DetailClient.Req) (*Model.Client, error) {

	detail, err := Repository.ClientRepository{}.Detail(req.ID)

	if err != nil {
		return detail, err
	}
	if detail.ID == 0 {
		return detail, errors.New("数据不存在或者已被删除。")
	}

	return detail, err
}
