package user_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/User/BindClient"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func UserBindClient(req *BindClient.Req) error {

	if err := validateBindClient(req.ID); err != nil {
		return err
	}

	var userRole []Model.UserClient

	for _, clientID := range req.ClientIDs {
		userRole = append(userRole, Model.UserClient{
			UserID:   req.ID,
			ClientID: clientID,
		})
	}

	return Repository.UserClientRepository{}.BatchStore(&userRole)
}

func validateBindClient(userID uint) error {
	var _ *Model.UserAuth
	var err error
	_, err = Repository.UserRepository{}.Detail(userID)
	if err != nil {
		return err
	}
	return nil
}
