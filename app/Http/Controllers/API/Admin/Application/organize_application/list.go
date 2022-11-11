package organize_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ListOrganize"
	"user_center/app/Repository"
)

func List(req *ListOrganize.Req) ([]ListOrganize.Resp, int, error) {

	list, total, err := Repository.OrganizeRepository{}.List(req)
	res := ListOrganize.Item(list)
	return res, total, err
}
