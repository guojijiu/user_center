package department_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/DetailDepartment"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Detail(req *DetailDepartment.Req) (*Model.Department, error) {

	detail, err := Repository.DepartmentRepository{}.Detail(req.ID)

	if err != nil {
		return detail, err
	}
	if detail.ID == 0 {
		return detail, errors.New("数据不存在或者已被删除。")
	}

	return detail, err
}
