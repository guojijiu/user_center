
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
	if db.Def().HasTable(Model.User{}.TableName()) {
		err = fmt.Errorf("uc_user_auth table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户登录鉴权表'").
		CreateTable(&Model.User{}).Error
	return
}

func (CreateUserTable) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.User{}).Error
	return
}
