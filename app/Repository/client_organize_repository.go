package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/GetBindOrganize"
	"user_center/app/Model"
)

type ClientOrganizeRepository struct{}

func (ClientOrganizeRepository) BatchStore(model *[]Model.ClientOrganize) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var organizeIDs []uint
		userID := (*model)[0].ClientID
		if err := tx.Model(&Model.ClientOrganize{}).Select("organize_id").Where("client_id = ?", userID).Find(&organizeIDs).Error; err != nil {
			return err
		}
		if organizeIDs != nil {
			if err := tx.Where("client_id = ?", userID).Delete(&model).Error; err != nil {
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

func (ClientOrganizeRepository) GetBindOrganizeIDs(clientID uint) []uint {
	var result []uint
	if err := DB.Model(&Model.ClientOrganize{}).Select("organize_id").Where("client_id = ?", clientID).Find(&result).Error; err != nil {
		return nil
	}

	return result
}

func (ClientOrganizeRepository) GetBindOrganize(clientID uint) ([]GetBindOrganize.Result, error) {
	var result []GetBindOrganize.Result
	if err := DB.Table("uc_organize as o").Joins("left join uc_mapping_client_organize as mco on o.id = mco.organize_id").
		Select("o.id,o.name").
		Where("mco.client_id = ?", clientID).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (ClientOrganizeRepository) DeleteByClientID(clientID uint) error {
	if err := DB.Where("client_id = ?", clientID).Delete(&Model.ClientOrganize{}).Error; err != nil {
		return err
	}
	return nil
}
