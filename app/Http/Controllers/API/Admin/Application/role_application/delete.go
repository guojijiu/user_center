package role_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DeleteRole"
	"user_center/app/Repository"
)

func Delete(req *DeleteRole.Req) error {

	detail, err := Repository.RoleRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}
	baseErr := Repository.RoleRepository{}.Delete(req.ID)
	if baseErr != nil {
		return baseErr
	}

	return err
}
