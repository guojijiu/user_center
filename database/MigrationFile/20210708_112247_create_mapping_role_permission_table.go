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
	if db.Def().HasTable(Model.RolePermission{}.TableName()) {
		err = fmt.Errorf("mapping_role_permission table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='角色权限关联表'").
		CreateTable(&Model.RolePermission{}).Error
	return
}

func (CreateMappingRolePermissionTable) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.RolePermission{}).Error
	return
}
