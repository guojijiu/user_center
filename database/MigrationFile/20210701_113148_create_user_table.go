package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateUserTable struct {
}

func (CreateUserTable) Key() string {
	return "20210701_113148_create_user_table.go"
}

func (CreateUserTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.User{}.TableName()) {
		err = fmt.Errorf("uc_user_auth table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户登录鉴权表'").
		Migrator().
		CreateTable(&Model.User{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateUserTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.User{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
