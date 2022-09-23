package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/permission_application"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/DeletePermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/DetailPermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/GetTreePermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/ListPermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/StorePermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/UpdatePermission"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

type PermissionController struct {
}

func (PermissionController) Store(c *gin.Context) {
	var err error
	var req StorePermission.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := permission_application.Store(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add permission fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (PermissionController) Update(c *gin.Context) {
	var err error
	var req UpdatePermission.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = permission_application.Update(&req, c)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update permission fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (PermissionController) Detail(c *gin.Context) {
	var err error
	var req DetailPermission.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := permission_application.Detail(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "detail permission fail", err), nil)
		return
	}

	Responses.Success(c, "success", DetailPermission.Item(res))
}

func (PermissionController) GetList(c *gin.Context) {
	var err error
	var req ListPermission.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, total, err := permission_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list permission fail", err), nil)
		return
	}

	body := map[string]interface{}{
		"data":  data,
		"total": total,
	}

	Responses.Success(c, "success", body)
}

func (PermissionController) Delete(c *gin.Context) {
	var err error
	var req DeletePermission.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := permission_application.Delete(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "delete permission fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (PermissionController) GetTree(c *gin.Context) {
	var err error
	var req GetTreePermission.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := permission_application.GetTree(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "tree permission fail", err), nil)
		return
	}

	Responses.Success(c, "success", res)
}
