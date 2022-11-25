package user_test

import (
	"fmt"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Web/Context/Auth/Login"
	"user_center/app/Http/Controllers/API/Web/Context/Auth/Logout"
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

// go test -v test/Feature/Web/User/auth_test.go -test.run TestLogin
func TestLogin(t *testing.T) {
	resp := httptest.Post("/api/web/login", Login.Req{
		Account:        "123456",
		Passwd:         "123456",
		CaptchaCode:    "111111",
		CaptchaCodeKey: "111",
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Web/Auth/auth_test.go -test.run TestLogin
func TestLogout(t *testing.T) {
	resp := httptest.Post("/api/web/logout", Logout.Req{})
	fmt.Println(resp.Body)
}
