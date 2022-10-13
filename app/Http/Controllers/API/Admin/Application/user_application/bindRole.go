package user_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/User/BindRole"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func UserBindRole(req *BindRole.Req) error {

	if err := validateBindRole(req.ID); err != nil {
		return err
	}

	var userRole []Model.UserRole

	for _, roleID := range req.RoleIDs {
		userRole = append(userRole, Model.UserRole{
			UserID: req.ID,
			RoleID: roleID,
		})
	}

	return Repository.UserRoleRepository{}.BatchStore(&userRole)
}

func validateBindRole(userID uint) error {
	var _ *Model.UserAuth
	var err error
	_, err = Repository.UserRepository{}.Detail(userID)
	if err != nil {
		return err
	}
	return nil
}
