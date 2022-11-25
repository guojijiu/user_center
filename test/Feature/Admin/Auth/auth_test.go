package user_test

import (
	"fmt"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Auth/Logout"
	"user_center/app/Http/Controllers/API/Web/Context/Auth/Login"
	"user_center/boot"
	"user_center/pkg/test"
)

var httptest *test.Http

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	httptest = test.New(app.GetEngineRouter())
	m.Run()
}

// go test -v test/Feature/Admin/Auth/auth_test.go -test.run TestLogin
func TestLogin(t *testing.T) {
	resp := httptest.Post("/api/admin/login", Login.Req{
		Account:        "aaa",
		Passwd:         "123456",
		CaptchaCode:    "111111",
		CaptchaCodeKey: "111",
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Auth/auth_test.go -test.run TestLogin
func TestLogout(t *testing.T) {
	header := test.Header{
		Key:   "token",
		Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjksImFjY291bnQiOiJhYWEiLCJleHAiOjE2NjkzODk2NTEsImlzcyI6IlVTRVJfQ0VOVEVSX1dFQiIsIm5iZiI6MTY2OTM2NzA1MX0.iDnesfmdHEMRvtxEGM3kcyMByjoRqjHg_4XgMsf-VoA",
	}
	resp := httptest.Post("/api/admin/logout", Logout.Req{}, header)
	fmt.Println(resp.Body)
}
