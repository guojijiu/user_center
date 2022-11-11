package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/department_application"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/DetailDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ForbiddenDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ListDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/StoreDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/UpdateDepartment"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

type DepartmentController struct {
}

func (DepartmentController) Store(c *gin.Context) {
	var err error
	var req StoreDepartment.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	//fmt.Printf("%+v", tool.JSONString(&req))

	storeErr := department_application.Store(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add department fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (DepartmentController) Update(c *gin.Context) {
	var err error
	var req UpdateDepartment.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = department_application.Update(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update department fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (DepartmentController) Detail(c *gin.Context) {
	var err error
	var req DetailDepartment.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := department_application.Detail(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "detail department fail", err), nil)
		return
	}

	Responses.Success(c, "success", DetailDepartment.Item(res))
}

func (DepartmentController) GetList(c *gin.Context) {
	var err error
	var req ListDepartment.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, total, err := department_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list department fail", err), nil)
		return
	}

	body := map[string]interface{}{
		"data":  data,
		"total": total,
	}

	Responses.Success(c, "success", body)
}

func (DepartmentController) Forbidden(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req ForbiddenDepartment.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := department_application.Forbidden(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "forbidden department fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}
