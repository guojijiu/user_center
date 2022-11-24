package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Web/Context/User/Register"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/encryption"
	"user_center/pkg/randc"
)

func RegisterUser(req *Register.Req) error {

	if validateErr := validateReq(req); validateErr != nil {
		return validateErr
	}

	pwd, PassErr := encryption.BcryptPasswdHash(req.Passwd)
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

	return Repository.UserRepository{}.Store(&user, &userInfo)
}

func validateReq(req *Register.Req) error {
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
