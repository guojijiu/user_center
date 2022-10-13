package Admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Application/client_application"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/DetailClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ForbiddenClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ListClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/StoreClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/UpdateClient"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/pkg/glog"
)

type ClientController struct {
}

func (ClientController) Store(c *gin.Context) {
	var err error
	var req StoreClient.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	storeErr := client_application.Store(&req)

	if storeErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "add client fail", storeErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (ClientController) Update(c *gin.Context) {
	var err error
	var req UpdateClient.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	err = client_application.Update(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "update client fail", err), nil)
		return
	}

	Responses.Success(c, "success", nil)
}

func (ClientController) Detail(c *gin.Context) {
	var err error
	var req DetailClient.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	res, err := client_application.Detail(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "detail client fail", err), nil)
		return
	}

	Responses.Success(c, "success", DetailClient.Item(res))
}

func (ClientController) GetList(c *gin.Context) {
	var err error
	var req ListClient.Req
	if err = c.ShouldBindQuery(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	data, total, err := client_application.List(&req)

	if err != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "list client fail", err), nil)
		return
	}

	body := map[string]interface{}{
		"data":  data,
		"total": total,
	}

	Responses.Success(c, "success", body)
}

func (ClientController) Forbidden(c *gin.Context) {
	var err error
	// 参赛不能为bool，值为false的情况会认为不存在
	var req ForbiddenClient.Req
	if err = c.ShouldBindJSON(&req); err != nil {
		glog.Default().Println("err=", err.Error())
		Responses.BadReq(c, err)
		return
	}

	forbiddenErr := client_application.Forbidden(&req)

	if forbiddenErr != nil {
		Responses.Failed(c, fmt.Sprintf("%s %s", "forbidden client fail", forbiddenErr), nil)
		return
	}

	Responses.Success(c, "success", nil)
}
