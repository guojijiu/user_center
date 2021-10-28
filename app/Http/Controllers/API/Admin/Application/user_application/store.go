package user_application

import (
	"errors"
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/User/StoreUser"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/db"
	"user_center/pkg/encryption"
	"user_center/pkg/randc"
	"user_center/pkg/tool"
)

func Store(req *StoreUser.Req) error {

	if validateErr := validateReq(req); validateErr != nil {
		return validateErr
	}

	pwd, PassErr := encryption.BcryptHash(req.Passwd)
	if PassErr != nil {
		return PassErr
	}

	var user Model.UserAuth
	user.UUID = randc.UUID()
	user.Account = req.Account
	user.Email = req.Email
	user.Phone = req.Phone
	user.Passwd = pwd

	var userInfo Model.UserInformation
	if req.Nickname != "" {
		userInfo.Nickname = req.Nickname
	}
	if req.HeaderImgPath != "" {
		userInfo.HeaderImgPath = req.HeaderImgPath
	}
	if req.Sex != "" {
		userInfo.Sex = req.Sex
	}
	if req.Birthday != "" {
		birthday := tool.TimeStrToDatetime(req.Birthday)
		userInfo.Birthday = &birthday
	}
	if req.Address != "" {
		userInfo.Address = req.Address
	}
	if req.Organization != "" {
		userInfo.Organization = req.Organization
	}
	if req.PersonalProfile != "" {
		userInfo.PersonalProfile = req.PersonalProfile
	}

	return db.Def().Transaction(func(tx *gorm.DB) error {
		baseErr := Repository.UserRepository{
			DB: tx,
		}.Store(&user)
		if baseErr != nil {
			return baseErr
		}
		userInfo.UserID = user.ID
		infoErr := Repository.UserInfoRepository{
			DB: tx,
		}.Store(&userInfo)
		if infoErr != nil {
			return infoErr
		}
		return nil
	})
}

func validateReq(req *StoreUser.Req) error {
	userByAccount, errByAccount := Repository.UserRepository{}.FindByAccount(req.Account)
	if errByAccount != nil {
		return errByAccount
	}
	if userByAccount.ID != 0 {
		return errors.New("账号已存在。")
	}

	userByEmail, errByEmail := Repository.UserRepository{}.FindByEmail(req.Email)
	if errByEmail != nil {
		return errByEmail
	}
	if userByEmail.ID != 0 {
		return errors.New("邮箱已存在。")
	}

	userByPhone, errByPhone := Repository.UserRepository{}.FindByPhone(req.Phone)
	if errByPhone != nil {
		return errByPhone
	}
	if userByPhone.ID != 0 {
		return errors.New("手机号已存在。")
	}
	return nil
}
