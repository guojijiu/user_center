package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateMappingUserClientTable struct {
}

func (CreateMappingUserClientTable) Key() string {
	return "20210708_103806_create_mapping_user_client_table.go"
}

func (CreateMappingUserClientTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.UserClient{}.TableName()) {
		err = fmt.Errorf("mapping_user_client table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='用户客户端关联表'").
		Migrator().
		CreateTable(&Model.UserClient{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateMappingUserClientTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.UserClient{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
