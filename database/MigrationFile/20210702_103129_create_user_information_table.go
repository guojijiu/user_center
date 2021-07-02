
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type Create_user_information_table struct {
}

func (Create_user_information_table) Key() string {
	return "20210702_103129_create_user_information_table.go"
}

func (Create_user_information_table) Up() (err error) {
	if db.Def().HasTable(Model.UserInformation{}.TableName()) {
		err = fmt.Errorf("uc_user_information table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户信息表'").
		CreateTable(&Model.UserInformation{}).Error
	return
}

func (Create_user_information_table) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.UserInformation{}).Error
	return
}
