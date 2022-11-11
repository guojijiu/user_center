package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateOrganizeTable struct {
}

func (CreateOrganizeTable) Key() string {
	return "20221110_105030_create_organize_table.go"
}

func (CreateOrganizeTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.Organize{}.TableName()) {
		err = fmt.Errorf("uc_organize table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='组织表'").
		Migrator().
		CreateTable(&Model.Organize{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
	}
	return
}

func (CreateOrganizeTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.Organize{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
