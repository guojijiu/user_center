package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Web/Context/User/ResetPasswd"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/encryption"
)

func Reset(req *ResetPasswd.Req) error {

	user, err := Repository.UserRepository{}.Detail(req.ID)
	if err != nil {
		return err
	}
	if validateErr := validateReset(user, req); validateErr != nil {
		return validateErr
	}
	var userInfo Model.UserInformation

	pwd, PassErr := encryption.BcryptPasswdHash(req.Passwd)
	if PassErr != nil {
		return PassErr
	}
	user.Passwd = pwd

	return Repository.UserRepository{}.Update(user, &userInfo, user.ID)
}

func validateReset(user *Model.UserAuth, req *ResetPasswd.Req) error {

	if user.ID != 0 {
		return errors.New("用户信息异常，请刷新页面重新重置密码")
	}
	if user.Account != req.Account {
		return errors.New("账号信息异常，请刷新页面重新重置密码。")
	}
	if req.Passwd != req.PasswdTwo {
		return errors.New("两次输入密码不一致，请检查后重新输入。")
	}
	return nil
}
