package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Model"
)

type RolePermissionRepository struct {
	DB *gorm.DB
}

func (RolePermissionRepository) BatchStore(rolePermission *[]Model.RolePermission) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var permissionIDs []uint
		roleID := (*rolePermission)[0].RoleID
		if err := tx.Model(&Model.RolePermission{}).Select("permission_id").Where("role_id = ?", roleID).Find(&permissionIDs).Error; err != nil {
			return err
		}
		if permissionIDs != nil {
			if err := tx.Where("role_id = ?", roleID).Delete(&rolePermission).Error; err != nil {
				return err
			}
		}
		if err := DB.Create(&rolePermission).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (RolePermissionRepository) GetBindPermissionIDs(roleID uint) []uint {
	var result []uint
	if err := DB.Model(&Model.RolePermission{}).Select("permission_id").Where("role_id = ?", roleID).Find(&result).Error; err != nil {
		return nil
	}

	return result
}

func (RolePermissionRepository) DeleteByRoleID(roleID uint) error {
	var rolePermission Model.RolePermission
	if err := DB.Where("role_id = ?", roleID).Delete(&rolePermission).Error; err != nil {
		return err
	}
	return nil
}
