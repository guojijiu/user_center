package user_test

import (
	"fmt"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Web/Context/User/GetCaptcha"
	"user_center/app/Http/Controllers/API/Web/Context/User/SendRegisterCode"
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

// go test -v test/Feature/Web/User/user_test.go -test.run TestSendRegisterCode
func TestSendRegisterCode(t *testing.T) {
	resp := httptest.Post("/api/web/send_register_code", SendRegisterCode.Req{
		Type: 1,
		//Phone: "13012113456",
		Email: "644522319@qq.com",
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Web/User/user_test.go -test.run TestGetCaptcha
func TestGetCaptcha(t *testing.T) {
	resp := httptest.Get("/api/web/captcha", GetCaptcha.Req{})
	fmt.Println(resp.Body)
}
