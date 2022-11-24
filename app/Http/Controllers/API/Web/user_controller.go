package Web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Web/Application/user_application"
	"user_center/app/Http/Controllers/API/Web/Context/User/DetailByForget"
	"user_center/app/Http/Controllers/API/Web/Context/User/GetCaptcha"
	"user_center/app/Http/Controllers/API/Web/Context/User/Register"
	"user_center/app/Http/Controllers/API/Web/Context/User/ResetPasswd"
	"user_center/app/Http/Controllers/API/Web/Context/User/SendForgetCode"
	"user_center/app/Http/Controllers/API/Web/Context/User/SendRegisterCode"
	"user_center/app/Http/Controllers/API/Web/Responses"
	"user_center/pkg/glog"
)

type UserController struct {
}

func (UserController) GetCaptcha(c *gin.Context) {
	var err error
	var req GetCaptcha.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	str, sendErr := user_application.Captcha(&req)

	if sendErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "get captcha code fail", sendErr), nil)
		return
	}

	Responses.Success(c, "success", str)
}

func (UserController) SendRegisterCode(c *gin.Context) {
	var err error
	var req SendRegisterCode.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	sendErr := user_application.SendRegister(&req)

	if sendErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "send register code fail", sendErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) Register(c *gin.Context) {
	var err error
	var req Register.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	sendErr := user_application.RegisterUser(&req)

	if sendErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "register fail", sendErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) SendForgetCode(c *gin.Context) {
	var err error
	var req SendForgetCode.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	sendErr := user_application.SendForget(&req)

	if sendErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "send forget code fail", sendErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) DetailByForget(c *gin.Context) {
	var err error
	var req DetailByForget.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	user, detailErr := user_application.GetDetailByForget(&req)

	if detailErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "get forget data fail", detailErr), nil)
		return
	}

	Responses.Success(c, "success", DetailByForget.Item(user))
}

func (UserController) ResetPasswd(c *gin.Context) {
	var err error
	var req ResetPasswd.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	resetErr := user_application.Reset(&req)

	if resetErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "reset passwd fail", resetErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}
