package user_application

import (
	"errors"
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/User/UpdateUser"
	"user_center/app/Model"
	"user_center/app/Repository/user_info_repository"
	"user_center/app/Repository/user_repository"
	"user_center/pkg/db"
	"user_center/pkg/tool"
)

func Update(req *UpdateUser.Req) error {

	detail, err := user_repository.UserRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}
	userInfoDetail, err := user_info_repository.UserInfoRepository{}.FindByUserID(req.ID)
	if err != nil {
		return nil
	}
	var user Model.UserAuth
	var userInfo Model.UserInformation

	if req.Account != "" && req.Account != detail.Account {
		user.Account = req.Account
	}
	if req.Phone != "" && req.Phone != detail.Phone {
		user.Phone = req.Phone
	}
	if req.Email != "" && req.Email != detail.Email {
		user.Email = req.Email
	}
	if req.Nickname != "" && req.Nickname != userInfoDetail.Nickname {
		userInfo.Nickname = req.Nickname
	}
	if req.HeaderImgPath != "" && req.HeaderImgPath != userInfoDetail.HeaderImgPath {
		userInfo.HeaderImgPath = req.HeaderImgPath
	}
	if req.Sex != "" && req.Sex != userInfoDetail.Sex {
		userInfo.Sex = req.Sex
	}
	if req.Birthday != "" {
		birthday := tool.TimeStrToDatetime(req.Birthday)
		userInfo.Birthday = &birthday
	}
	if req.Address != "" && req.Address != userInfoDetail.Address {
		userInfo.Address = req.Address
	}
	if req.Organization != "" && req.Organization != userInfoDetail.Organization {
		userInfo.Organization = req.Organization
	}
	if req.PersonalProfile != "" && req.PersonalProfile != userInfoDetail.PersonalProfile {
		userInfo.PersonalProfile = req.PersonalProfile
	}

	return db.Def().Transaction(func(tx *gorm.DB) error {
		baseErr := user_repository.UserRepository{
			DB: tx,
		}.Update(&user, req.ID)
		if baseErr != nil {
			return baseErr
		}
		infoError := user_info_repository.UserInfoRepository{
			DB: tx,
		}.Update(&userInfo, userInfoDetail.ID)
		if infoError != nil {
			return infoError
		}
		return nil
	})
}
