package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateUserInformationTable struct {
}

func (CreateUserInformationTable) Key() string {
	return "20210702_103129_create_user_information_table.go"
}

func (CreateUserInformationTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.UserInformation{}.TableName()) {
		err = fmt.Errorf("uc_user_information table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户信息表'").
		Migrator().
		CreateTable(&Model.UserInformation{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateUserInformationTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.UserInformation{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
