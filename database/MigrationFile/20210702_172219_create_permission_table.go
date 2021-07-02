
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type Create_permission_table struct {
}

func (Create_permission_table) Key() string {
	return "20210702_172219_create_permission_table.go"
}

func (Create_permission_table) Up() (err error) {
	if db.Def().HasTable(Model.Permission{}.TableName()) {
		err = fmt.Errorf("uc_permission table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='权限表'").
		CreateTable(&Model.Permission{}).Error
	return
}

func (Create_permission_table) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.Permission{}).Error
	return
}
