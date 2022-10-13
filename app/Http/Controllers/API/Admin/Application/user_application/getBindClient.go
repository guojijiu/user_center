package user_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindClient"
	"user_center/app/Repository"
)

func GetUserBindClient(req *GetBindClient.Req) ([]GetBindClient.Result, error) {
	return Repository.UserClientRepository{}.GetClientByUserID(req.ID)
}
