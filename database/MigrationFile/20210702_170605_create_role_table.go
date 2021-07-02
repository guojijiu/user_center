
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type Create_role_table struct {
}

func (Create_role_table) Key() string {
	return "20210702_170605_create_role_table.go"
}

func (Create_role_table) Up() (err error) {
	if db.Def().HasTable(Model.Role{}.TableName()) {
		err = fmt.Errorf("uc_role table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='角色表'").
		CreateTable(&Model.Role{}).Error
	return
}

func (Create_role_table) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.Role{}).Error
	return
}
