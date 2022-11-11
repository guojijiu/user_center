package client_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Client/BindOrganize"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func ClientBindOrganize(req *BindOrganize.Req) error {

	if err := validateBindOrganize(req.ID); err != nil {
		return err
	}

	var clientOrganize []Model.ClientOrganize

	for _, organizeID := range req.OrganizeIDs {
		clientOrganize = append(clientOrganize, Model.ClientOrganize{
			ClientID:   req.ID,
			OrganizeID: organizeID,
		})
	}

	return Repository.ClientOrganizeRepository{}.BatchStore(&clientOrganize)
}

func validateBindOrganize(userID uint) error {
	var _ *Model.Client
	var err error
	_, err = Repository.ClientRepository{}.Detail(userID)
	if err != nil {
		return err
	}
	return nil
}
