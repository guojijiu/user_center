package role_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/StoreRole"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Store(req *StoreRole.Req) error {

	if validateErr := validateReq(req); validateErr != nil {
		return validateErr
	}

	var roleModel Model.Role

	roleModel.ClientID = req.ClientID
	roleModel.Name = req.Name
	roleModel.Mark = req.Mark
	roleModel.Remark = req.Remark
	roleModel.Sort = req.Sort

	baseErr := Repository.RoleRepository{}.Store(&roleModel)
	if baseErr != nil {
		return baseErr
	}

	return nil
}

func validateReq(req *StoreRole.Req) error {
	role, err := Repository.RoleRepository{}.FindByMark(req.Mark)
	if err != nil {
		return err
	}
	if role.ID != 0 {
		return errors.New("角色mark已存在。")
	}

	return nil
}
