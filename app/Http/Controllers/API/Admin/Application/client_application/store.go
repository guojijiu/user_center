package client_application

import (
	"errors"
	"fmt"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/StoreClient"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/randc"
)

func Store(req *StoreClient.Req) error {

	if validateErr := validateStore(req); validateErr != nil {
		return validateErr
	}

	var client Model.Client
	client.UUID = randc.UUID()
	client.Name = req.Name
	client.Mark = req.Mark

	if req.Remark != "" {
		client.Remark = req.Remark
	}

	return Repository.ClientRepository{}.Store(&client)
}

func validateStore(req *StoreClient.Req) error {
	client, err := Repository.ClientRepository{}.FindByMark(req.Mark)
	if err != nil {
		return err
	}
	if client.ID != 0 {
		return errors.New(fmt.Sprintf("唯一标识符：%s，已存在。", req.Mark))
	}

	return nil
}
