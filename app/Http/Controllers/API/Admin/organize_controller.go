package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/organize_application"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/DetailOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ForbiddenOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ListOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/StoreOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/UpdateOrganize"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

type OrganizeController struct {
}

func (OrganizeController) Store(c *gin.Context) {
	var err error
	var req StoreOrganize.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	//fmt.Printf("%+v", tool.JSONString(&req))

	storeErr := organize_application.Store(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add organize fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (OrganizeController) Update(c *gin.Context) {
	var err error
	var req UpdateOrganize.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = organize_application.Update(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update organize fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (OrganizeController) Detail(c *gin.Context) {
	var err error
	var req DetailOrganize.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := organize_application.Detail(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "detail organize fail", err), nil)
		return
	}

	Responses.Success(c, "success", DetailOrganize.Item(res))
}

func (OrganizeController) GetList(c *gin.Context) {
	var err error
	var req ListOrganize.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, total, err := organize_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list organize fail", err), nil)
		return
	}

	body := map[string]interface{}{
		"data":  data,
		"total": total,
	}

	Responses.Success(c, "success", body)
}

func (OrganizeController) Forbidden(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req ForbiddenOrganize.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := organize_application.Forbidden(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "forbidden organize fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}
