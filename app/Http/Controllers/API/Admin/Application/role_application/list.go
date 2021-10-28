package role_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Role/ListRole"
	"user_center/app/Repository"
)

func List(req *ListRole.Req) ([]ListRole.Resp, int, error) {

	list, total, err := Repository.RoleRepository{}.List(req)
	res := ListRole.Item(list)
	return res, total, err
}
