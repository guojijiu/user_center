package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateMappingClientOrganizeTable struct {
}

func (CreateMappingClientOrganizeTable) Key() string {
	return "20221110_131816_create_mapping_client_organize_table.go"
}

func (CreateMappingClientOrganizeTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.ClientOrganize{}.TableName()) {
		err = fmt.Errorf("mapping_client_organizess table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='客户端组织关联表'").
		Migrator().
		CreateTable(&Model.ClientOrganize{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateMappingClientOrganizeTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.ClientOrganize{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
