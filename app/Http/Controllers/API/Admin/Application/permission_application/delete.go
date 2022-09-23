package permission_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/DeletePermission"
	"user_center/app/Repository"
)

func Delete(req *DeletePermission.Req) error {

	detail, err := Repository.PermissionRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}
	baseErr := Repository.PermissionRepository{}.Delete(req.ID)
	if baseErr != nil {
		return baseErr
	}

	return err
}
