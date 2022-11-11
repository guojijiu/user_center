package department_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ForbiddenDepartment"
	"user_center/app/Repository"
)

func Forbidden(req *ForbiddenDepartment.Req) error {

	detail, err := Repository.DepartmentRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	if req.IsForbidden == 1 {
		return Repository.DepartmentRepository{}.Forbidden(req.ID)
	} else {
		return Repository.DepartmentRepository{}.UnForbidden(req.ID)
	}
}
