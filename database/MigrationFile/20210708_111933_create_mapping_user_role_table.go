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
	if db.Def().Migrator().HasTable(Model.UserRole{}.TableName()) {
		err = fmt.Errorf("mapping_user_role table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户角色关联表'").
		Migrator().
		CreateTable(&Model.UserRole{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateMappingUserRoleTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.UserRole{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
