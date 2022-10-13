package client_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ForbiddenClient"
	"user_center/app/Repository"
)

func Forbidden(req *ForbiddenClient.Req) error {

	detail, err := Repository.ClientRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	if req.IsForbidden == 1 {
		return Repository.ClientRepository{}.Forbidden(req.ID)
	} else {
		return Repository.ClientRepository{}.UnForbidden(req.ID)
	}
}
