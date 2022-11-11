package Repository

import (
	"gorm.io/gorm"
	"time"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ListOrganize"
	"user_center/app/Model"
	"user_center/pkg/tool"
)

type OrganizeRepository struct{}

func (OrganizeRepository) Store(storeModel *Model.Organize) error {
	err := DB.Create(&storeModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (OrganizeRepository) Detail(id uint) (*Model.Organize, error) {
	var model Model.Organize
	err := DB.Where("id = ? AND deleted_at is null", id).First(&model).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &model, nil
}

func (OrganizeRepository) Update(updateModel *Model.Organize, ID uint) error {
	if err := DB.Model(&Model.Organize{}).Where("id = ?", ID).Updates(updateModel).Error; err != nil {
		return err
	}
	return nil
}

func (OrganizeRepository) List(req *ListOrganize.Req) ([]Model.Organize, int, error) {
	var (
		listModel []Model.Organize
		total     int64
	)
	offset, limit := tool.PageCoverLimit(req.Page, req.Size)
	query := DB.Model(&Model.Organize{})
	query.Count(&total)
	if err := query.Offset(offset).Limit(limit).Find(&listModel).Error; err != nil {
		return listModel, 0, err
	}
	return listModel, int(total), nil
}

func (OrganizeRepository) Forbidden(ID uint) error {
	nowTime := tool.TimeStrToDatetime(time.Now().Format("2006-01-02 15:04:05"))
	if err := DB.Model(&Model.Organize{}).
		Where("id = ?", ID).
		Update("forbade_at", &nowTime).Error; err != nil {
		return err
	}
	return nil
}

func (OrganizeRepository) UnForbidden(ID uint) error {
	if err := DB.Model(&Model.Organize{}).
		Where("id = ?", ID).
		Update("forbade_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (OrganizeRepository) FindByMark(mark string) (*Model.Organize, error) {

	var model Model.Organize
	err := DB.Where("mark = ? AND deleted_at is null", mark).First(&model).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &model, nil
}
