
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreatePermissionTable struct {
}

func (CreatePermissionTable) Key() string {
	return "20210702_172219_create_permission_table.go"
}

func (CreatePermissionTable) Up() (err error) {
	if db.Def().HasTable(Model.Permission{}.TableName()) {
		err = fmt.Errorf("uc_permission table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='权限表'").
		CreateTable(&Model.Permission{}).Error
	return
}

func (CreatePermissionTable) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.Permission{}).Error
	return
}
