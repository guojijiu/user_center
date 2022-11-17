package Web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Web/Application/user_application"
	"user_center/app/Http/Controllers/API/Web/Context/User/GetCaptcha"
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
