package user_info_repository

import (
	"gorm.io/gorm"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type UserInfoRepository struct {
	DB *gorm.DB
}

// 默认db库选择
var DB = db.Def()

func (UserInfoRepository) Store(userInfo *Model.UserInformation) error {
	err := DB.Create(&userInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func (UserInfoRepository) Detail(id uint) (*Model.UserInformation, error) {
	var userInfo Model.UserInformation
	err := DB.Where("id = ? AND deleted_at is null", id).First(&userInfo).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &userInfo, nil
}

func (UserInfoRepository) FindByUserID(userID uint) (*Model.UserInformation, error) {
	var userInfo Model.UserInformation
	err := DB.Where("user_id = ? AND deleted_at is null", userID).First(&userInfo).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &userInfo, nil
}

func (UserInfoRepository) Update(userInfoUpdate *Model.UserInformation, ID uint) error {
	var userInfo Model.UserInformation
	if err := DB.Model(&userInfo).Where("id = ?", ID).Updates(&userInfoUpdate).Error; err != nil {
		return err
	}
	return nil
}
