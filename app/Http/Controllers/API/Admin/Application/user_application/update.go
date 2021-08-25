package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/UpdateUser"
	"user_center/app/Model"
	"user_center/app/Repository/user_repository"
)

func Update(req *UpdateUser.Req) error {

	detail, err := user_repository.UserRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}
	var user Model.UserAuth

	if req.Account != "" && req.Account != detail.Account {
		user.Account = req.Account
	}
	if req.Phone != "" && req.Phone != detail.Phone {
		user.Phone = req.Phone
	}
	if req.Email != "" && req.Email != detail.Email {
		user.Email = req.Email
	}

	return user_repository.UserRepository{}.Update(&user)
}
