package organize_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ForbiddenOrganize"
	"user_center/app/Repository"
)

func Forbidden(req *ForbiddenOrganize.Req) error {

	detail, err := Repository.OrganizeRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	if req.IsForbidden == 1 {
		return Repository.OrganizeRepository{}.Forbidden(req.ID)
	} else {
		return Repository.OrganizeRepository{}.UnForbidden(req.ID)
	}
}
