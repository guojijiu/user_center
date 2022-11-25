package Admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/auth_application"
	"user_center/app/Http/Controllers/API/Admin/Context/Auth/Login"
	"user_center/app/Http/Controllers/API/Admin/Context/Auth/Logout"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
	jwtauth "user_center/pkg/jwt"
)

type AuthController struct {
}

func (AuthController) Login(c *gin.Context) {
	var err error
	var req Login.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	result, loginErr := auth_application.LoginUser(&req)
	if loginErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "login fail", loginErr), nil)
		return
	}

	Responses.Success(c, "success", result)
}

func (AuthController) Logout(c *gin.Context) {
	var err error
	var req Logout.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	var newData jwtauth.CustomClaims
	jwtData, _ := c.Get("claims")
	resByre, _ := json.Marshal(jwtData)
	_ = json.Unmarshal(resByre, &newData)

	logoutErr := auth_application.LogoutUser(newData.ID)
	if logoutErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "logout fail", logoutErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}
