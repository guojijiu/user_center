
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type Create_client_table struct {
}

func (Create_client_table) Key() string {
	return "20210702_155243_create_client_table.go"
}

func (Create_client_table) Up() (err error) {
	if db.Def().HasTable(Model.Client{}.TableName()) {
		err = fmt.Errorf("uc_client table alreay exist")
		return
	}
	err = db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='客户端表'").
		CreateTable(&Model.Client{}).Error
	return
}

func (Create_client_table) Down() (err error) {
	err = db.Def().DropTableIfExists(&Model.Client{}).Error
	return
}
