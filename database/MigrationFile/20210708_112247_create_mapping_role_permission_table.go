package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateMappingRolePermissionTable struct {
}

func (CreateMappingRolePermissionTable) Key() string {
	return "20210708_112247_create_mapping_role_permission_table.go"
}

func (CreateMappingRolePermissionTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.RolePermission{}.TableName()) {
		err = fmt.Errorf("mapping_role_permission table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='角色权限关联表'").
		Migrator().
		CreateTable(&Model.RolePermission{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateMappingRolePermissionTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.RolePermission{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
