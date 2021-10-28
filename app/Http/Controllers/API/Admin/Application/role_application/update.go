package role_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/UpdateRole"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Update(req *UpdateRole.Req) error {

	detail, err := Repository.RoleRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	var roleModel Model.Role

	if req.Name != "" && req.Name != detail.Name {
		roleModel.Name = req.Name
	}
	if req.Mark != "" && req.Mark != detail.Mark {
		roleModel.Mark = req.Mark
	}
	if req.Sort != 0 && req.Sort != detail.Sort {
		roleModel.Sort = req.Sort
	}
	if req.Remark != "" && req.Remark != detail.Remark {
		roleModel.Remark = req.Remark
	}

	UpdateErr := Repository.RoleRepository{}.Update(&roleModel, req.ID)
	if UpdateErr != nil {
		return UpdateErr
	}
	return nil
}
