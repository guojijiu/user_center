package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateUsersTableMigrate struct {
}

func (CreateUsersTableMigrate) Key() string {
	return "2020_5_7_17_59_create_users_table"
}

func (CreateUsersTableMigrate) Up() (err error) {
	if db.Def().HasTable(Model.User{}.TableName()) {
		err = fmt.Errorf("users table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户登录鉴权'").
		CreateTable(&Model.User{}).Error
	return
}

func (CreateUsersTableMigrate) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.User{}).Error
	return
}
