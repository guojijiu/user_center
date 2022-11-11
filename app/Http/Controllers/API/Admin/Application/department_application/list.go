package department_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ListDepartment"
	"user_center/app/Repository"
)

func List(req *ListDepartment.Req) ([]ListDepartment.Resp, int, error) {

	list, total, err := Repository.DepartmentRepository{}.List(req)
	res := ListDepartment.Item(list)
	return res, total, err
}
