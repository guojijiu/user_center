package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/GetBindDepartment"
	"user_center/app/Model"
)

type RoleDepartmentRepository struct{}

func (RoleDepartmentRepository) BatchStore(model *[]Model.RoleDepartment) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var departmentIDs []uint
		userID := (*model)[0].RoleID
		if err := tx.Model(&Model.RoleDepartment{}).Select("department_id").Where("role_id = ?", userID).Find(&departmentIDs).Error; err != nil {
			return err
		}
		if departmentIDs != nil {
			if err := tx.Where("role_id = ?", userID).Delete(&model).Error; err != nil {
				return err
			}
		}
		if err := tx.Create(&model).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (RoleDepartmentRepository) GetBindDepartmentIDs(roleID uint) []uint {
	var result []uint
	if err := DB.Model(&Model.RoleDepartment{}).Select("department_id").Where("role_id = ?", roleID).Find(&result).Error; err != nil {
		return nil
	}

	return result
}

func (RoleDepartmentRepository) GetBindDepartment(roleID uint) ([]GetBindDepartment.Result, error) {
	var result []GetBindDepartment.Result
	if err := DB.Table("uc_department as d").Joins("left join uc_mapping_role_department as mrd on d.id = mrd.department_id").
		Select("d.id,d.name").
		Where("mrd.role_id = ?", roleID).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (RoleDepartmentRepository) DeleteByRoleID(roleID uint) error {
	if err := DB.Where("role_id = ?", roleID).Delete(&Model.RoleDepartment{}).Error; err != nil {
		return err
	}
	return nil
}
