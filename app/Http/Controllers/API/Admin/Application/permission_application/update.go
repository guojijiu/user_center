package permission_application

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/UpdatePermission"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/tool"
)

func Update(req *UpdatePermission.Req, httpRequest *gin.Context) error {

	detail, err := Repository.PermissionRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	if req.Remark != "" {
		detailByMark, _ := Repository.PermissionRepository{}.FindByMark(req.Mark)
		if detailByMark.ID != 0 {
			return errors.New(fmt.Sprintf("参数mark：%s，已存在，请重新添加。", req.Mark))
		}
	}

	var permissionModel Model.Permission

	if req.Name != "" && req.Name != detail.Name {
		permissionModel.Name = req.Name
	}
	if req.Mark != "" && req.Mark != detail.Mark {
		permissionModel.Mark = req.Mark
	}
	if req.Sort != 0 && req.Sort != detail.Sort {
		permissionModel.Sort = req.Sort
	}
	if req.Remark != "" && req.Remark != detail.Remark {
		permissionModel.Remark = req.Remark
	}
	if req.Type != 0 && req.Remark != detail.Remark {
		permissionModel.Type = req.Type
	}
	if req.ParentID != 0 && req.ParentID != detail.ParentID {
		permissionModel.ParentID = req.ParentID
	}
	if iconPath, isOk := httpRequest.GetPostForm(req.IconPath); isOk {
		permissionModel.IconPath = iconPath
	}
	if routeName, isOk := httpRequest.GetPostForm(req.RouteName); isOk {
		permissionModel.RouteName = routeName
	}
	if routePath, isOk := httpRequest.GetPostForm(req.RoutePath); isOk {
		permissionModel.RoutePath = routePath
	}
	if modulePath, isOk := httpRequest.GetPostForm(req.ModulePath); isOk {
		permissionModel.ModulePath = modulePath
	}
	if requestMethod, isOk := httpRequest.GetPostForm(req.RequestMethod); isOk {
		permissionModel.RequestMethod = requestMethod
	}
	if hiddenAt, isOk := httpRequest.GetPostForm(req.HiddenAt); isOk {
		permissionModel.HiddenAt = tool.TimeStrToDatetime(hiddenAt)
	}

	UpdateErr := Repository.PermissionRepository{}.Update(&permissionModel, req.ID)
	if UpdateErr != nil {
		return UpdateErr
	}
	return nil
}
