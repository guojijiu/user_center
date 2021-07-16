package user_repository

import (
	"github.com/jinzhu/gorm"
	"time"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ForbiddenUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Model"
	"user_center/pkg/db"
	"user_center/pkg/tool"
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

func (UserRepository) Update(user *Model.UserAuth) (bool, error) {
	if err := db.Def().Updates(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (UserRepository) List(req *ListUser.Req) ([]*Model.UserAuth, error) {
	var userList []*Model.UserAuth
	if err := db.Def().Limit(req.Size).Offset(req.Page).Find(&userList).Error; err != nil {
		return userList, err
	}
	return userList, nil
}

func (UserRepository) Forbidden(req ForbiddenUser.Req) (bool, error) {
	nowTime := tool.TimeStrToDatetime(time.Now().Format("2006-01-02 15:04:05"))
	if err := db.Def().Model(&Model.UserAuth{}).
		Where("id = ?", req.ID).
		Update("forbade_at", &nowTime).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (UserRepository) UnForbidden(req ForbiddenUser.Req) (bool, error) {
	if err := db.Def().Model(&Model.UserAuth{}).
		Where("id = ?", req.ID).
		Update("forbade_at", nil).Error; err != nil {
		return false, err
	}
	return true, nil
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
