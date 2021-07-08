package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/user_application"
	"user_center/app/Http/Controllers/API/Admin/Context/User/StoreUser"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

func Store(c *gin.Context) {
	var err error
	var req StoreUser.StoreReq
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, storeErr := user_application.StoreUserApp(&req)

	if res == false && storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add user fail", err), nil)
		return
	}

	Responses.Success(c, "success", res)
}
