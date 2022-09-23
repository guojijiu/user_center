package migrator

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"user_center/command"
	"user_center/pkg/db"
	migrate2 "user_center/pkg/migrate"
)

var CMDmigrator = &command.Command{
	UsageLine: "migrator [command]",
	Short:     "执行数据库迁移",
	Long:      usageDoc(),
	Run:       RunMigration,
}

// go run main.go migrator
func init() {
	command.CMD.Register(CMDmigrator)
}

// RunMigration 执行数据库迁移管理
func RunMigration(cmd *command.Command, args []string) int {
	var err error
	if len(args) == 0 {
		cmd.Info("Running all migration")
		err = migrate(cmd)
	} else {
		err = cmd.Flag.Parse(args[1:])
		if err != nil {
			_ = fmt.Errorf("Parse migrate options error: %s \n", err.Error())
		}
		subcmd := args[0]
		switch subcmd {
		case "migrate":
			err = migrate(cmd)
			break
		case "rollback":
			err = rollback(cmd)
			break
		case "fresh":
			err = fresh(cmd)
			break
		case "delete":
			err = del(cmd)
			break
		case "tables":
			tables, err := tables(cmd)
			if err == nil {
				for _, table := range tables {
					cmd.Info(table)
				}
			}
			break
		default:
			usage()
		}
		if err != nil {
			cmd.Error(err.Error())
		}
	}

	return 0
}

func usageDoc() string {
	return `migrater verson: migrator/1.0.0 
           Usage: ./uims migrate/rollback`
}

func usage() {
	_, _ = fmt.Fprintf(os.Stdout, `migrater verson: migrater/1.0.0
Usage: migrater migrate/rollback/tables/fresh/delete

Options:
`)
	flag.PrintDefaults()
}

// 获取所有表
func tables(cmd *command.Command) ([]string, error) {
	tables := []string{}
	err := db.Def().Raw("show tables").Pluck("Tables_in_mysql", &tables).Error
	if err != nil {
		return tables, err
	}
	return tables, nil
}

// 删除所有表
func del(cmd *command.Command) error {
	tables, err := tables(cmd)
	if err != nil {
		return err
	}
	for _, table := range tables {
		cmd.Info("[deleting] " + table)
		err = db.Def().Exec("drop table " + table).Error
		if err != nil {
			return err
		}
		cmd.Info("[deleted] " + table + " success")
	}
	return nil
}

// 删除所有表, 并重新执行迁移
func fresh(cmd *command.Command) error {
	// 删除所有表
	cmd.Info("start delete tables")
	err := del(cmd)
	if err != nil {
		return errors.Wrap(err, "delete table failed")
	}
	cmd.Info(">>>> migrating tables")
	// 执行迁移
	err = migrate(cmd)
	if err != nil {
		return errors.Wrap(err, "migrate failed")
	}
	cmd.Info(">>>> migrated tables")
	return nil
}

// 执行迁移
func migrate(cmd *command.Command) error {
	var err error
	err = db.Def().AutoMigrate(&migrate2.Migration{})
	if err != nil {
		return errors.Wrap(err, "create migrate table failed")
	}
	mfs := migrate2.GetNeedMigrateFiles(migrate2.MigrateFiles)
	nextBatch := migrate2.GetNextBatchNo()
	cmd.Info(fmt.Sprintf("migrate file has %d", len(mfs)))
	if len(mfs) == 0 {
		cmd.Info("No migrate file need migration")
		return nil
	}
	for _, mf := range mfs {
		cmd.Info(fmt.Sprintf("[migrating] %s ...", mf.Key()))
		err = mf.Up()
		if err != nil {
			return errors.Wrapf(err, "[migrate failed] %s: %s", mf.Key(), err.Error())
		}
		err = migrate2.CreateMigrate(mf.Key(), nextBatch)
		if err != nil {
			return errors.Wrapf(err, "[migrate failed] %s: %s", mf.Key(), err.Error())
		}
		cmd.Info(fmt.Sprintf("[migrated] %s successed", mf.Key()))
	}
	return nil
}

// 回滚操作
func rollback(cmd *command.Command) error {
	var err error
	mfs := migrate2.GetNeedRollbackKeys(1)
	cmd.Info(fmt.Sprintf("Rollback file has %d \n", len(mfs)))
	if len(mfs) == 0 {
		cmd.Info("No migrate file need rollback")
		return nil
	}
	for _, mf := range mfs {
		cmd.Info(fmt.Sprintf("[Rollbacking] %s ... \n", mf.Key()))
		err = mf.Down()
		if err != nil {
			return errors.Wrapf(err, "[Rollback failed] %s: %s", mf.Key(), err.Error())
		}
		err = migrate2.DeleteMigrate(mf.Key())
		if err != nil {
			return errors.Wrapf(err, "[Rollback failed] %s: %s", mf.Key(), err.Error())
		}
		cmd.Info(fmt.Sprintf("[Rollbacked] %s successed \n", mf.Key()))
	}
	return nil
}
