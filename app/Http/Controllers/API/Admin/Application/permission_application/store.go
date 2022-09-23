package permission_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/StorePermission"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/tool"
)

func Store(req *StorePermission.Req) error {

	if validateErr := validateReq(req); validateErr != nil {
		return validateErr
	}

	var permissionModel Model.Permission

	permissionModel.ClientID = req.ClientID
	permissionModel.Name = req.Name
	permissionModel.Mark = req.Mark
	permissionModel.Remark = req.Remark
	permissionModel.Sort = req.Sort
	permissionModel.Type = req.Type
	if req.ParentID != 0 {
		permissionModel.ParentID = req.ParentID
	}
	if req.Remark != "" {
		permissionModel.Remark = req.Remark
	}
	if req.IconPath != "" {
		permissionModel.IconPath = req.IconPath
	}
	if req.RouteName != "" {
		permissionModel.RouteName = req.RouteName
	}
	if req.RoutePath != "" {
		permissionModel.RoutePath = req.RoutePath
	}
	if req.ModulePath != "" {
		permissionModel.ModulePath = req.ModulePath
	}
	if req.RequestMethod != "" {
		permissionModel.RequestMethod = req.RequestMethod
	}
	if req.HiddenAt != "" {
		permissionModel.HiddenAt = tool.TimeStrToDatetime(req.HiddenAt)
	}

	baseErr := Repository.PermissionRepository{}.Store(&permissionModel)
	if baseErr != nil {
		return baseErr
	}

	return nil
}

func validateReq(req *StorePermission.Req) error {
	role, err := Repository.PermissionRepository{}.FindByMark(req.Mark)
	if err != nil {
		return err
	}
	if role.ID != 0 {
		return errors.New("权限mark已存在。")
	}

	return nil
}
