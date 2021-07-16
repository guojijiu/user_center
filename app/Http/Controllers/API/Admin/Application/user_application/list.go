package user_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Model"
	"user_center/app/Repository/user_repository"
)

func List(req *ListUser.Req) ([]*Model.UserAuth, error) {

	res, err := user_repository.UserRepository{}.List(req)
	return res, err
}
