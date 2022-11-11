package Repository

import (
	"gorm.io/gorm"
	"time"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ListDepartment"
	"user_center/app/Model"
	"user_center/pkg/tool"
)

type DepartmentRepository struct{}

func (DepartmentRepository) Store(storeModel *Model.Department) error {
	err := DB.Create(&storeModel).Error
	if err != nil {
		return err
	}
	return nil
}

func (DepartmentRepository) Detail(id uint) (*Model.Department, error) {
	var model Model.Department
	err := DB.Where("id = ? AND deleted_at is null", id).First(&model).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &model, nil
}

func (DepartmentRepository) Update(updateModel *Model.Department, ID uint) error {
	if err := DB.Model(&Model.Department{}).Where("id = ?", ID).Updates(updateModel).Error; err != nil {
		return err
	}
	return nil
}

func (DepartmentRepository) List(req *ListDepartment.Req) ([]Model.Department, int, error) {
	var (
		listModel []Model.Department
		total     int64
	)
	offset, limit := tool.PageCoverLimit(req.Page, req.Size)
	query := DB.Model(&Model.Department{})
	query.Count(&total)
	if err := query.Offset(offset).Limit(limit).Find(&listModel).Error; err != nil {
		return listModel, 0, err
	}
	return listModel, int(total), nil
}

func (DepartmentRepository) Forbidden(ID uint) error {
	nowTime := tool.TimeStrToDatetime(time.Now().Format("2006-01-02 15:04:05"))
	if err := DB.Model(&Model.Department{}).
		Where("id = ?", ID).
		Update("forbade_at", &nowTime).Error; err != nil {
		return err
	}
	return nil
}

func (DepartmentRepository) UnForbidden(ID uint) error {
	if err := DB.Model(&Model.Department{}).
		Where("id = ?", ID).
		Update("forbade_at", nil).Error; err != nil {
		return err
	}
	return nil
}

func (DepartmentRepository) FindByMark(mark string) (*Model.Department, error) {

	var model Model.Department
	err := DB.Where("mark = ? AND deleted_at is null", mark).First(&model).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &model, nil
}
