package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type CreateDepartmentTable struct {
}

func (CreateDepartmentTable) Key() string {
	return "20221111_144737_create_department_table.go"
}

func (CreateDepartmentTable) Up() (err error) {
	if db.Def().Migrator().HasTable(Model.Department{}.TableName()) {
		err = fmt.Errorf("department table alreay exist")
		return
	}
	if createErr := db.Def().
		Set("gorm:table_options", "CHARSET=utf8mb4,COMMENT='部门表'").
		Migrator().
		CreateTable(&Model.Department{}); createErr != nil {
		_ = fmt.Errorf(createErr.Error())
		return
	}
	return
}

func (CreateDepartmentTable) Down() (err error) {
	if dropErr := db.Def().Migrator().DropTable(&Model.Department{}); dropErr != nil {
		_ = fmt.Errorf(dropErr.Error())
	}
	return
}
