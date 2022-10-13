package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/User/GetBindClient"
	"user_center/app/Model"
)

type UserClientRepository struct {
	DB *gorm.DB
}

func (UserClientRepository) BatchStore(userClient *[]Model.UserClient) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		var clientIDs []uint
		userID := (*userClient)[0].UserID
		if err := tx.Model(&Model.UserClient{}).Select("client_id").Where("user_id = ?", userID).Find(&clientIDs).Error; err != nil {
			return err
		}
		if clientIDs != nil {
			if err := tx.Where("user_id = ?", userID).Delete(&userClient).Error; err != nil {
				return err
			}
		}
		if err := DB.Create(&userClient).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func (UserClientRepository) GetClientByUserID(userID uint) ([]GetBindClient.Result, error) {
	var result []GetBindClient.Result
	if err := DB.Table("uc_client as c").Joins("left join uc_mapping_user_client as muc on c.id = muc.client_id").
		Select("c.id,c.name").
		Where("muc.user_id = ?", userID).
		Scan(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
