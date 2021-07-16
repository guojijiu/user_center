package user_application

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"user_center/app/Http/Controllers/API/Admin/Context/User/StoreUser"
	"user_center/app/Model"
	"user_center/app/Repository/user_repository"
	"user_center/pkg/encryption"
)

func Store(req *StoreUser.Req) (bool, error) {

	userByAccount, errByAccount := user_repository.UserRepository{}.FindByAccount(req.Account)
	if errByAccount != nil {
		return false, errByAccount
	}
	if userByAccount.ID != 0 {
		return false, errors.New("账号已存在。")
	}

	userByEmail, errByEmail := user_repository.UserRepository{}.FindByEmail(req.Email)
	if errByEmail != nil {
		return false, errByEmail
	}
	if userByEmail.ID != 0 {
		return false, errors.New("邮箱已存在。")
	}

	userByPhone, errByPhone := user_repository.UserRepository{}.FindByPhone(req.Phone)
	if errByPhone != nil {
		return false, errByPhone
	}
	if userByPhone.ID != 0 {
		return false, errors.New("手机号已存在。")
	}

	pwd, PassErr := encryption.BcryptHash(req.Passwd)
	if PassErr != nil {
		return false, PassErr
	}
	var user Model.UserAuth

	user.UUID = uuid.NewV1().String()
	user.Account = req.Account
	user.Email = req.Email
	user.Phone = req.Phone
	user.Passwd = pwd
	return user_repository.UserRepository{}.Store(&user)
}
