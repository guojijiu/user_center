
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateMappingUserRoleTable struct {
}

func (CreateMappingUserRoleTable) Key() string {
	return "20210708_111933_create_mapping_user_role_table.go"
}

func (CreateMappingUserRoleTable) Up() (err error) {
	if db.Def().HasTable(Model.UserRole{}.TableName()) {
		err = fmt.Errorf("mapping_user_role table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户角色关联表'").
		CreateTable(&Model.UserRole{}).Error
	return
}

func (CreateMappingUserRoleTable) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.UserRole{}).Error
	return
}
