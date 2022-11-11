package client_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Client/GetBindOrganize"
	"user_center/app/Repository"
)

func GetUserBindOrganize(req *GetBindOrganize.Req) ([]GetBindOrganize.Result, error) {
	return Repository.ClientOrganizeRepository{}.GetBindOrganize(req.ID)
}
