package auth_application

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user_center/app/Http/Controllers/API/Admin/Context/Auth/Login"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func LoginUser(req *Login.Req) (Login.Resp, error) {

	user, validateErr := validateLogin(req)
	if validateErr != nil {
		return Login.Resp{}, validateErr
	}
	return Login.Item(user), nil
}

func validateLogin(req *Login.Req) (*Model.UserAuth, error) {

	var userByAccount *Model.UserAuth
	userByAccount, errByAccount := Repository.UserRepository{}.FindByAccount(req.Account)
	if errByAccount != nil {
		return userByAccount, errByAccount
	}
	if userByAccount.ID == 0 {
		return userByAccount, errors.New("账号或者密码错误，请检查后重新输入。")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userByAccount.Passwd), []byte(req.Passwd)); err != nil {
		return userByAccount, errors.New("账号或者密码错误，请检查后重新输入。")
	}

	return userByAccount, nil
}
