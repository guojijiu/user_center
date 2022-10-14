package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/ListPermission"
	"user_center/app/Model"
	"user_center/pkg/tool"
)

type PermissionRepository struct{}

func (PermissionRepository) Store(permission *Model.Permission) error {
	err := DB.Create(&permission).Error
	if err != nil {
		return err
	}
	return nil
}

func (PermissionRepository) Detail(id uint) (*Model.Permission, error) {
	var permission Model.Permission
	err := DB.Where("id = ? AND deleted_at is null", id).First(&permission).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return &permission, err
	}
	return &permission, nil
}

func (PermissionRepository) Update(permissionUpdate *Model.Permission, ID uint) error {
	var permission Model.Permission
	if err := DB.Model(&permission).Where("id = ? and deleted_at is null", ID).Updates(permissionUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (PermissionRepository) List(req *ListPermission.Req) ([]Model.Permission, int, error) {
	var (
		permissionList []Model.Permission
		total          int64
	)
	offset, limit := tool.PageCoverLimit(req.Page, req.Size)
	query := DB.Model(&Model.Permission{})
	query.Count(&total)
	if err := query.Offset(offset).Limit(limit).Find(&permissionList).Error; err != nil {
		return permissionList, 0, err
	}
	return permissionList, int(total), nil
}

func (PermissionRepository) FindByMark(mark string) (*Model.Permission, error) {

	var permission Model.Permission
	err := DB.Where("mark = ? AND deleted_at is null", mark).First(&permission).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &permission, nil
}

func (PermissionRepository) Delete(ID uint) error {
	var permission Model.Permission
	err := DB.Where("id = ? AND deleted_at is null", ID).Delete(&permission).Error
	if err != nil {
		return err
	}
	return nil
}

func (PermissionRepository) GetAllByClientID(clientID uint) ([]Model.Permission, error) {
	var permissionList []Model.Permission
	if err := DB.Model(&Model.Permission{}).Where("client_id = ? AND deleted_at is null", clientID).Find(&permissionList).Error; err != nil {
		return permissionList, err
	}
	return permissionList, nil
}
