package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ForbiddenUser"
	"user_center/app/Repository"
)

func Forbidden(req *ForbiddenUser.Req) error {

	detail, err := Repository.UserRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	if req.IsForbidden == 1 {
		return Repository.UserRepository{}.Forbidden(req.ID)
	} else {
		return Repository.UserRepository{}.UnForbidden(req.ID)
	}
}
