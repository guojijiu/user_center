package permission_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/GetTreePermission"
	"user_center/app/Repository"
)

func GetTree(req *GetTreePermission.Req) ([]GetTreePermission.Resp, error) {

	data, err := Repository.PermissionRepository{}.GetAllByClientID(req.ClientID)

	filterData := GetTreePermission.Item(data)
	return filterData, err
}
