package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/user_application"
	"user_center/app/Http/Controllers/API/Admin/Context/User/BindClient"
	"user_center/app/Http/Controllers/API/Admin/Context/User/BindRole"
	"user_center/app/Http/Controllers/API/Admin/Context/User/DetailUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ExportUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ForbiddenUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindClient"
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindPermission"
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindRole"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/StoreUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/UpdateUser"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

type UserController struct {
}

func (UserController) Store(c *gin.Context) {
	var err error
	var req StoreUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	//fmt.Printf("%+v", tool.JSONString(&req))

	storeErr := user_application.Store(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add user fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) Update(c *gin.Context) {
	var err error
	var req UpdateUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = user_application.Update(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update user fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) Detail(c *gin.Context) {
	var err error
	var req DetailUser.Req
	if err = c.ShouldBindQuery(&req); err != nil {
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

func (UserController) GetList(c *gin.Context) {
	var err error
	var req ListUser.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, total, err := user_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list user fail", err), nil)
		return
	}

	body := map[string]interface{}{
		"data":  data,
		"total": total,
	}

	Responses.Success(c, "success", body)
}

func (UserController) Forbidden(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req ForbiddenUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := user_application.Forbidden(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "forbidden user fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) BindRole(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req BindRole.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = user_application.UserBindRole(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "user bind role fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) GetBindRole(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req GetBindRole.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, err := user_application.GetUserBindRole(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "get user bind role fail", err), nil)
		return
	}

	Responses.Success(c, "success", data)
}

func (UserController) GetBindPermission(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req GetBindPermission.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, err := user_application.GetUserBindPermission(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "get user permission fail", err), nil)
		return
	}

	Responses.Success(c, "success", data)
}

func (UserController) BindClient(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req BindClient.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = user_application.UserBindClient(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "user bind client fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (UserController) GetBindClient(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req GetBindClient.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, err := user_application.GetUserBindClient(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "get user bind client fail", err), nil)
		return
	}

	Responses.Success(c, "success", data)
}

func (UserController) ExportUser(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req ExportUser.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	filePath, err := user_application.ExportUserData(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "export user data fail", err), nil)
		return
	}

	Responses.Success(c, "success", filePath)
}
