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
	if db.Def().Migrator().HasTable(Model.Role{}.TableName()) {
		err = fmt.Errorf("uc_role table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='角色表'").
		Migrator().
		CreateTable(&Model.Role{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
	}
	return
}

func (CreateRoleTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.Role{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
