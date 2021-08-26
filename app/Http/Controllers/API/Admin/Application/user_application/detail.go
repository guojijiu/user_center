package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/DetailUser"
	"user_center/app/Model"
	"user_center/app/Repository/user_repository"
)

func Detail(req *DetailUser.Req) (*Model.UserAuth, error) {

	detail, err := user_repository.UserRepository{}.DetailOfAll(req.ID)

	if err != nil {
		return detail, err
	}
	if detail.ID == 0 {
		return detail, errors.New("数据不存在或者已被删除。")
	}

	return detail, err
}
