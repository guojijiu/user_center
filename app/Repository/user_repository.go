package Repository

import (
	"gorm.io/gorm"
	"time"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ExportUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Model"
	"user_center/pkg/db"
	"user_center/pkg/tool"
)

type UserRepository struct{}

func (UserRepository) Store(user *Model.UserAuth, userInfo *Model.UserInformation) error {
	return db.Def().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		userInfo.UserID = user.ID
		if err := tx.Create(&userInfo).Error; err != nil {
			return err
		}
		return nil
	})
}

func (UserRepository) Detail(id uint) (*Model.UserAuth, error) {
	var user Model.UserAuth
	err := DB.Where("id = ? AND deleted_at is null", id).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) Update(userUpdate *Model.UserAuth, userInfoUpdate *Model.UserInformation, ID uint) error {
	return db.Def().Transaction(func(tx *gorm.DB) error {
		var user Model.UserAuth
		if err := DB.Model(&user).Where("id = ?", ID).Updates(userUpdate).Error; err != nil {
			return err
		}
		var userInfo Model.UserInformation
		if err := DB.Model(&userInfo).Where("user_id = ?", ID).Updates(&userInfoUpdate).Error; err != nil {
			return err
		}
		return nil
	})

}

func (UserRepository) List(req *ListUser.Req) ([]Model.UserAuth, int, error) {
	var (
		userList []Model.UserAuth
		total    int64
	)
	offset, limit := tool.PageCoverLimit(req.Page, req.Size)
	query := DB.Model(&Model.UserAuth{})
	query.Count(&total)
	if err := query.Preload("UserInfo").Offset(offset).Limit(limit).Find(&userList).Error; err != nil {
		return userList, 0, err
	}
	return userList, int(total), nil
}

func (UserRepository) Forbidden(userID uint) error {
	nowTime := tool.TimeStrToDatetime(time.Now().Format("2006-01-02 15:04:05"))
	if err := DB.Model(&Model.UserAuth{}).
		Where("id = ?", userID).
		Update("forbade_at", &nowTime).Error; err != nil {
		return err
	}
	return nil
}

func (UserRepository) UnForbidden(userID uint) error {
	if err := DB.Model(&Model.UserAuth{}).
		Where("id = ?", userID).
		Update("forbade_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (UserRepository) FindByAccount(account string) (*Model.UserAuth, error) {

	var user Model.UserAuth
	err := DB.Where("account = ? AND deleted_at is null", account).First(&user).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByEmail(email string) (*Model.UserAuth, error) {

	var user Model.UserAuth
	err := DB.Where("email = ? AND deleted_at is null", email).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) FindByPhone(phone string) (*Model.UserAuth, error) {

	var user Model.UserAuth
	err := DB.Where("phone = ? AND deleted_at is null", phone).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) DetailOfAll(id uint) (*Model.UserAuth, error) {
	var user Model.UserAuth
	err := DB.Where("id = ? AND deleted_at is null", id).
		Preload("UserInfo").
		First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (UserRepository) GetDataByExport(req *ExportUser.Req) ([]ExportUser.Resp, error) {
	var result []ExportUser.Resp
	query := DB.Table("uc_user_auth as ua").
		Joins("left join uc_mapping_user_client as muc on ua.id = muc.user_id").
		Joins("left join uc_client as c on c.id = muc.client_id").
		Select("ua.id,ua.email,c.name as client_name,ua.created_at,ua.forbade_at")

	if req.ClientID != 0 {
		query.Where("muc.client_id = ?", req.ClientID)
	}

	if err := query.Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
