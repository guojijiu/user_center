package permission_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/ListPermission"
	"user_center/app/Repository"
)

func List(req *ListPermission.Req) ([]ListPermission.Resp, int, error) {

	list, total, err := Repository.PermissionRepository{}.List(req)
	res := ListPermission.Item(list)
	return res, total, err
}
