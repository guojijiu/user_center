package Repository

import (
	"gorm.io/gorm"
	"time"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ListClient"
	"user_center/app/Model"
	"user_center/pkg/tool"
)

type ClientRepository struct {
	DB *gorm.DB
}

func (ClientRepository) Store(client *Model.Client) error {
	err := DB.Create(&client).Error
	if err != nil {
		return err
	}
	return nil
}

func (ClientRepository) Detail(id uint) (*Model.Client, error) {
	var user Model.Client
	err := DB.Where("id = ? AND deleted_at is null", id).First(&user).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &user, nil
}

func (ClientRepository) Update(clientUpdate *Model.Client, ID uint) error {
	var client Model.Client
	if err := DB.Model(&client).Where("id = ?", ID).Updates(clientUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) List(req *ListClient.Req) ([]Model.Client, int, error) {
	var (
		clientList []Model.Client
		total      int64
	)
	offset, limit := tool.PageCoverLimit(req.Page, req.Size)
	query := DB.Model(&Model.Client{})
	query.Count(&total)
	if err := query.Offset(offset).Limit(limit).Find(&clientList).Error; err != nil {
		return clientList, 0, err
	}
	return clientList, int(total), nil
}

func (ClientRepository) Forbidden(clientID uint) error {
	nowTime := tool.TimeStrToDatetime(time.Now().Format("2006-01-02 15:04:05"))
	if err := DB.Model(&Model.Client{}).
		Where("id = ?", clientID).
		Update("forbade_at", &nowTime).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) UnForbidden(clientID uint) error {
	if err := DB.Model(&Model.Client{}).
		Where("id = ?", clientID).
		Update("forbade_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (ClientRepository) FindByMark(mark string) (*Model.Client, error) {

	var client Model.Client
	err := DB.Where("mark = ? AND deleted_at is null", mark).First(&client).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &client, nil
}
