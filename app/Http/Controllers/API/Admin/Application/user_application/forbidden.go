package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ForbiddenUser"
	"user_center/app/Repository/user_repository"
)

func Forbidden(req *ForbiddenUser.Req) error {

	detail, err := user_repository.UserRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	if req.IsForbidden == true {
		return user_repository.UserRepository{}.Forbidden(*req)
	} else {
		return user_repository.UserRepository{}.UnForbidden(*req)
	}
}
