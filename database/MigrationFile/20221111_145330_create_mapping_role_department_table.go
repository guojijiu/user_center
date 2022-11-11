package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateMappingRoleDepartmentTable struct {
}

func (CreateMappingRoleDepartmentTable) Key() string {
	return "20221111_145330_create_mapping_role_department_table.go"
}

func (CreateMappingRoleDepartmentTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.RoleDepartment{}.TableName()) {
		err = fmt.Errorf("mapping_role_department table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='角色部门关联表'").
		Migrator().
		CreateTable(&Model.RoleDepartment{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateMappingRoleDepartmentTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.RoleDepartment{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
