
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
	if db.Def().Migrator().HasTable(Model.Permission{}.TableName()) {
		err = fmt.Errorf("uc_permission table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='权限表'").
		Migrator().
		CreateTable(&Model.Permission{});createErr !=nil{
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreatePermissionTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.Permission{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
