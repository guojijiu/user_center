package user_test

import (
	"fmt"
	genid "github.com/srlemon/gen-id"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Web/Context/User/DetailByForget"
	"user_center/app/Http/Controllers/API/Web/Context/User/GetCaptcha"
	"user_center/app/Http/Controllers/API/Web/Context/User/Register"
	"user_center/app/Http/Controllers/API/Web/Context/User/ResetPasswd"
	"user_center/app/Http/Controllers/API/Web/Context/User/SendForgetCode"
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

// go test -v test/Feature/Web/User/user_test.go -test.run TestRegister
func TestRegister(t *testing.T) {
	resp := httptest.Post("/api/web/register", Register.Req{
		Type:        1,
		CaptchaCode: "123456",
		VerifyCode:  "123456",
		Account:     genid.NewGeneratorData().Name,
		Phone:       genid.NewGeneratorData().PhoneNum,
		Email:       genid.NewGeneratorData().Email,
		Passwd:      "123456",
		Nickname:    genid.NewGeneratorData().GeneratorName(),
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Web/User/user_test.go -test.run TestSendForgetCode
func TestSendForgetCode(t *testing.T) {
	resp := httptest.Post("/api/web/send_forget_code", SendForgetCode.Req{
		Type:  1,
		Phone: genid.NewGeneratorData().PhoneNum,
		Email: genid.NewGeneratorData().Email,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Web/User/user_test.go -test.run TestDetailByForget
func TestDetailByForget(t *testing.T) {
	resp := httptest.Get("/api/web/detail_by_forget", DetailByForget.Req{
		ForgetType: "dasdasd",
		ForgetData: "dfdfgdfgdf",
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Web/User/user_test.go -test.run TestResetPasswd
func TestResetPasswd(t *testing.T) {
	resp := httptest.Post("/api/web/reset_passwd", ResetPasswd.Req{
		ID:        1,
		Account:   genid.NewGeneratorData().Name,
		Email:     genid.NewGeneratorData().Email,
		Phone:     genid.NewGeneratorData().PhoneNum,
		Passwd:    "123456",
		PasswdTwo: "123456",
	})
	fmt.Println(resp.Body)
}
