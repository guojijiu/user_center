package Repository

import (
	"gorm.io/gorm"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/ListRole"
	"user_center/app/Model"
	"user_center/pkg/tool"
)

type RoleRepository struct {
	DB *gorm.DB
}

func (RoleRepository) Store(role *Model.Role) error {
	err := DB.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (RoleRepository) Detail(id uint) (*Model.Role, error) {
	var role Model.Role
	err := DB.Where("id = ? AND deleted_at is null", id).First(&role).Error
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return &role, err
	}
	return &role, nil
}

func (RoleRepository) Update(roleUpdate *Model.Role, ID uint) error {
	var role Model.Role
	if err := DB.Model(&role).Where("id = ? and deleted_at is null", ID).Updates(roleUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (RoleRepository) List(req *ListRole.Req) ([]Model.Role, int, error) {
	var (
		roleList []Model.Role
		total    int64
	)
	offset, limit := tool.PageCoverLimit(req.Page, req.Size)
	query := DB.Model(&Model.Role{})
	query.Count(&total)
	if err := query.Offset(offset).Limit(limit).Find(&roleList).Error; err != nil {
		return roleList, 0, err
	}
	return roleList, int(total), nil
}

func (RoleRepository) FindByMark(mark string) (*Model.Role, error) {

	var role Model.Role
	err := DB.Where("mark = ? AND deleted_at is null", mark).First(&role).Error

	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return nil, err
	}
	return &role, nil
}

func (RoleRepository) Delete(ID uint) error {
	var role Model.Role
	err := DB.Where("id = ? AND deleted_at is null", ID).Delete(&role).Error
	if err != nil {
		return err
	}
	return nil
}
