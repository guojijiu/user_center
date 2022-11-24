package user_application

import (
	"user_center/app/Domain/Cache"
	"user_center/app/Http/Controllers/API/Web/Context/User/DetailByForget"
	"user_center/app/Model"
	"user_center/app/Repository"
	conf "user_center/config"
	"user_center/pkg/encryption"
)

func GetDetailByForget(req *DetailByForget.Req) (*Model.UserAuth, error) {
	var (
		user *Model.UserAuth
		err  error
	)
	codeCache := &Cache.CodeCache{}
	str, err := decryptStr(req.ForgetType)
	if err != nil {
		return user, err
	}
	_, err = codeCache.SetCacheKey("user:forget_code", str).Get()
	if err != nil {
		return user, err
	}

	forgetType, _ := encryptStr(req.ForgetType)
	switch forgetType {
	case "1":
		user, err = Repository.UserRepository{}.FindByEmail(str)
		break
	case "2":
		user, err = Repository.UserRepository{}.FindByPhone(str)
		break
	default:
		err = nil
	}
	return user, err
}

func decryptStr(val string) (string, error) {
	cipher, err := encryption.NewEncryptor(conf.APPKey, encryption.CIPHER_AES_256_CBC)
	if err != nil {
		return "", err
	}
	cryptStr := cipher.DecryptWithCBC(val)
	return cryptStr, nil
}
