package Jwt_auth_test

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"testing"
	"time"
	"user_center/boot"
	"user_center/pkg/jwt"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

// go test -v test/Feature/Pkg/Jwt/jwt_auth_test.go -test.run TestSetJwt
func TestSetJwt(t *testing.T) {
	// Create the Claims
	claims := jwtauth.CustomClaims{
		11,
		"aaa",
		jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,   // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600*6, // 过期时间 六小时
			Issuer:    "USER_CENTER_WEB",
		},
	}
	token, err := jwtauth.NewJWT().CreateToken(claims)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
}

// go test -v test/Feature/Pkg/Jwt/jwt_auth_test.go -test.run TestParseToken
func TestParseToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsImFjY291bnQiOiJhYWEiLCJleHAiOjE2NjkzNzg0NTAsImlzcyI6IlVTRVJfQ0VOVEVSX1dFQiIsIm5iZiI6MTY2OTM1NTg1MH0.eN3idsROmDMU9vbjj_q8mwoNvix-cqjcbmi-NZz0pO8"
	token, err := jwtauth.NewJWT().ParseToken(tokenStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
}

// go test -v test/Feature/Pkg/Jwt/jwt_auth_test.go -test.run TestRefreshToken
func TestRefreshToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjksImFjY291bnQiOiJhYWEiLCJleHAiOjE2NjkzODMwODksImlzcyI6IlVTRVJfQ0VOVEVSX1dFQiIsIm5iZiI6MTY2OTM2MDQ4OX0.1V7yz2vjeqw9kNosEQQ5TwK62vW1hMfnyh_THneoYdQ"
	token, err := jwtauth.NewJWT().RefreshToken(tokenStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
}
