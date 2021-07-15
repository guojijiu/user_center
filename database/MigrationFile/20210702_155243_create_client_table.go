package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateClientTable struct {
}

func (CreateClientTable) Key() string {
	return "20210702_155243_create_client_table.go"
}

func (CreateClientTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.Client{}.TableName()) {
		err = fmt.Errorf("uc_client table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='客户端表'").
		Migrator().
		CreateTable(&Model.Client{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
	}
	return
}

func (CreateClientTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.Client{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
