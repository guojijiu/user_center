package user_repository

import (
	"github.com/jinzhu/gorm"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type UserRepository struct {
}

func (UserRepository) Store(user *Model.UserAuth) (bool, error) {
	err := db.Def().Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (UserRepository) Detail(id uint) (*Model.UserAuth, error) {
	var user Model.UserAuth
	err := db.Def().Where("id = ? AND deleted_at is null", id).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByAccount(account string) (*Model.UserAuth, error) {

	var user Model.UserAuth
	err := db.Def().Where("account = ? AND deleted_at is null", account).First(&user).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByEmail(email string) (*Model.UserAuth, error) {

	var user Model.UserAuth
	err := db.Def().Where("email = ? AND deleted_at is null", email).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByPhone(phone string) (*Model.UserAuth, error) {

	var user Model.UserAuth
	err := db.Def().Where("phone = ? AND deleted_at is null", phone).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}
