package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/DetailUser"
	"user_center/app/Model"
	"user_center/app/Repository/user_repository"
)

func Detail(req *DetailUser.Req) (*Model.UserAuth, error) {

	detail, err := user_repository.UserRepository{}.Detail(req.ID)

	var user *Model.UserAuth
	if err != nil {
		return user, err
	}
	if detail.ID == 0 {
		return user, errors.New("数据不存在或者已被删除。")
	}

	res, err := user_repository.UserRepository{}.Detail(req.ID)
	return res, err
}
