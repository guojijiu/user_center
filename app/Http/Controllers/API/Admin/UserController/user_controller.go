package UserController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/user_application"
	"user_center/app/Http/Controllers/API/Admin/Context/User/DetailUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ForbiddenUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/StoreUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/UpdateUser"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

func Store(c *gin.Context) {
	var err error
	var req StoreUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, storeErr := user_application.Store(&req)

	if res == false && storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add user fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", res)
}

func Update(c *gin.Context) {
	var err error
	var req UpdateUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := user_application.Update(&req)

	if res == false && err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update user fail", err), nil)
		return
	}

	Responses.Success(c, "success", res)
}

func Detail(c *gin.Context) {
	var err error
	var req DetailUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := user_application.Detail(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "detail user fail", err), nil)
		return
	}

	Responses.Success(c, "success", DetailUser.Item(res))
}

func GetList(c *gin.Context) {
	var err error
	var req ListUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := user_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list user fail", err), nil)
		return
	}

	Responses.Success(c, "success", ListUser.Item(res))
}

func Forbidden(c *gin.Context) {
	var err error
	var req ForbiddenUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, storeErr := user_application.Forbidden(&req)

	if res == false && storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "forbidden user fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", res)
}
