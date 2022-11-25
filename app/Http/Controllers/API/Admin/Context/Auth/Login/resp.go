package Login

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
	"user_center/app/Domain/Cache"
	"user_center/app/Model"
	jwtauth "user_center/pkg/jwt"
)

type Resp struct {
	ID      uint   `comment:"主键ID" json:"id"`
	Account string `comment:"账号" json:"account"`
	Token   string `comment:"访问token" json:"token"`
}

func Item(model *Model.UserAuth) Resp {
	claims := jwtauth.CustomClaims{
		ID:      model.ID,
		Account: model.Account,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,   // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600*6, // 过期时间 六小时
			Issuer:    "USER_CENTER_WEB",
		},
	}
	token, _ := jwtauth.NewJWT().CreateToken(claims)
	cacheUser(model)
	return Resp{
		ID:      model.ID,
		Account: model.Account,
		Token:   token,
	}
}

func cacheUser(model *Model.UserAuth) {
	cache := Cache.UserCache{}
	userCache := Cache.UserInfo{
		ID:      model.ID,
		Account: model.Account,
		Phone:   model.Phone,
		Email:   model.Email,
	}
	_, _ = cache.SetCacheKey(model.ID).Store(&userCache)
}
