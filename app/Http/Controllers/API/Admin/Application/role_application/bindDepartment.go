package role_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Role/BindDepartment"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func RoleBindDepartment(req *BindDepartment.Req) error {

	if err := validateBindDepartment(req.ID); err != nil {
		return err
	}

	var roleDepartment []Model.RoleDepartment

	for _, departmentID := range req.DepartmentIDs {
		roleDepartment = append(roleDepartment, Model.RoleDepartment{
			RoleID:       req.ID,
			DepartmentID: departmentID,
		})
	}

	return Repository.RoleDepartmentRepository{}.BatchStore(&roleDepartment)
}

func validateBindDepartment(roleID uint) error {
	var _ *Model.Client
	var err error
	_, err = Repository.RoleRepository{}.Detail(roleID)
	if err != nil {
		return err
	}
	return nil
}
