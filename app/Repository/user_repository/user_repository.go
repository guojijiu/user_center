package user_repository

import (
	"github.com/jinzhu/gorm"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type UserRepository struct {
}

func (UserRepository) Store(user *Model.User) (bool, error) {
	err := db.Def().Create(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (UserRepository) Detail(id uint) (*Model.User, error) {
	var user Model.User
	err := db.Def().Where("id = ? AND deleted is null", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByAccount(account string) (*Model.User, error) {

	var user Model.User
	err := db.Def().Where("account = ? AND deleted is null", account).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByEmail(email string) (*Model.User, error) {

	var user Model.User
	err := db.Def().Where("email = ? AND deleted is null", email).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByPhone(phone string) (*Model.User, error) {

	var user Model.User
	err := db.Def().Where("phone = ? AND deleted is null", phone).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}
