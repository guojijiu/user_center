package role_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Role/BindPermission"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Bind(req *BindPermission.Req) error {

	if err := validate(req.ID); err != nil {
		return err
	}

	var rolePermissionModel []Model.RolePermission

	for _, permissionID := range req.PermissionIDs {
		rolePermissionModel = append(rolePermissionModel, Model.RolePermission{
			RoleID:       req.ID,
			PermissionID: permissionID,
		})
	}

	return Repository.RolePermissionRepository{}.BatchStore(&rolePermissionModel)
}

func validate(roleID uint) error {
	var _ *Model.Role
	var err error
	_, err = Repository.RoleRepository{}.Detail(roleID)
	if err != nil {
		return err
	}
	return nil
}
