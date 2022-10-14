package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindRole"
	"user_center/app/Model"
)

type UserRoleRepository struct{}

func (UserRoleRepository) BatchStore(userRole *[]Model.UserRole) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var roleIDs []uint
		userID := (*userRole)[0].UserID
		if err := tx.Model(&Model.UserRole{}).Select("role_id").Where("user_id = ?", userID).Find(&roleIDs).Error; err != nil {
			return err
		}
		if roleIDs != nil {
			if err := tx.Where("user_id = ?", userID).Delete(&userRole).Error; err != nil {
				return err
			}
		}
		if err := tx.Create(&userRole).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (UserRoleRepository) GetBindRoleIDs(userID uint) []uint {
	var result []uint
	if err := DB.Model(&Model.UserRole{}).Select("role_id").Where("user_id = ?", userID).Find(&result).Error; err != nil {
		return nil
	}

	return result
}

func (UserRoleRepository) GetBindRole(userID uint) ([]GetBindRole.Result, error) {
	var result []GetBindRole.Result
	if err := DB.Table("uc_role as r").Joins("left join uc_mapping_user_role as mur on r.id = mur.role_id").
		Select("r.id,r.name").
		Where("mur.user_id = ?", userID).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (UserRoleRepository) DeleteByUserID(userID uint) error {
	var userRole Model.UserRole
	if err := DB.Where("user_id = ?", userID).Delete(&userRole).Error; err != nil {
		return err
	}
	return nil
}
