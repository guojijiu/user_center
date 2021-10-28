package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/role_application"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DeleteRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DetailRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/ListRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/StoreRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/UpdateRole"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

type RoleController struct {
}

func (RoleController) Store(c *gin.Context) {
	var err error
	var req StoreRole.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := role_application.Store(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add role fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (RoleController) Update(c *gin.Context) {
	var err error
	var req UpdateRole.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = role_application.Update(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update role fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (RoleController) Detail(c *gin.Context) {
	var err error
	var req DetailRole.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := role_application.Detail(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "detail role fail", err), nil)
		return
	}

	Responses.Success(c, "success", DetailRole.Item(res))
}

func (RoleController) GetList(c *gin.Context) {
	var err error
	var req ListRole.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, total, err := role_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list role fail", err), nil)
		return
	}

	body := map[string]interface{}{
		"data":  data,
		"total": total,
	}

	Responses.Success(c, "success", body)
}

func (RoleController) Delete(c *gin.Context) {
	var err error
	var req DeleteRole.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := role_application.Delete(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "delete role fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}
