
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateRoleTable struct {
}

func (CreateRoleTable) Key() string {
	return "20210702_170605_create_role_table.go"
}

func (CreateRoleTable) Up() (err error) {
	if db.Def().HasTable(Model.Role{}.TableName()) {
		err = fmt.Errorf("uc_role table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='角色表'").
		CreateTable(&Model.Role{}).Error
	return
}

func (CreateRoleTable) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.Role{}).Error
	return
}
