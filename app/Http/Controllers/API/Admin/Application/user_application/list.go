package user_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Repository/user_repository"
)

func List(req *ListUser.Req) ([]ListUser.Resp, int, error) {

	list, total, err := user_repository.UserRepository{}.List(req)
	res := ListUser.Item(list)
	return res, total, err
}
