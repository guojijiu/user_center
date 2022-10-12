package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindPermission"
	"user_center/app/Repository"
)

func GetUserBindPermission(req *GetBindPermission.Req) ([]GetBindPermission.Result, error) {

	roleIDs := Repository.UserRoleRepository{}.GetBindRoleIDs(req.ID)
	if roleIDs == nil {
		return nil, errors.New("用户未绑定角色数据。")
	}
	return Repository.RolePermissionRepository{}.GetPermissionByRoleIDs(roleIDs)
}
