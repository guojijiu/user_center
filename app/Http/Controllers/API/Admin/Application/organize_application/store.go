package organize_application

import (
	"errors"
	"fmt"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/StoreOrganize"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/randc"
)

func Store(req *StoreOrganize.Req) error {

	if validateErr := validateStore(req); validateErr != nil {
		return validateErr
	}

	var model Model.Organize
	model.UUID = randc.UUID()
	model.Name = req.Name
	model.Mark = req.Mark

	if req.Remark != "" {
		model.Remark = req.Remark
	}

	return Repository.OrganizeRepository{}.Store(&model)
}

func validateStore(req *StoreOrganize.Req) error {
	client, err := Repository.OrganizeRepository{}.FindByMark(req.Mark)
	if err != nil {
		return err
	}
	if client.ID != 0 {
		return errors.New(fmt.Sprintf("唯一标识符：%s，已存在。", req.Mark))
	}

	return nil
}
