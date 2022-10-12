package user_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindRole"
	"user_center/app/Repository"
)

func GetUserBindRole(req *GetBindRole.Req) ([]GetBindRole.Result, error) {
	return Repository.UserRoleRepository{}.GetBindRole(req.ID)
}
