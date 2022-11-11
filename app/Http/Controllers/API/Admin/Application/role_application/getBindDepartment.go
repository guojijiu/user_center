package role_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Role/GetBindDepartment"
	"user_center/app/Repository"
)

func GetRoleBindDepartment(req *GetBindDepartment.Req) ([]GetBindDepartment.Result, error) {
	return Repository.RoleDepartmentRepository{}.GetBindDepartment(req.ID)
}
