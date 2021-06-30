package migrate_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user_center/boot"
	"user_center/pkg/db"
	"user_center/pkg/migrate"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestGetAllMigrations(t *testing.T) {
	allMigrations := migrate.GetAllMigrations()
	assert.NotNil(t, allMigrations)
}

func TestGetNeedMigrateFiles(t *testing.T) {
	err := db.Def().Exec("truncate migrate_file").Error
	assert.Nil(t, err)
	needMigrateFiles := migrate.GetNeedMigrateFiles(migrate.MigrateFiles)
	assert.NotNil(t, needMigrateFiles)
	assert.Equal(t, len(migrate.MigrateFiles), len(needMigrateFiles))
	m := migrate.Migration{
		Migration: migrate.MigrateFiles[0].Key(),
		Batch:     1,
	}
	db.Def().Save(&m)

	needMigrateFiles = migrate.GetNeedMigrateFiles(migrate.MigrateFiles)
	assert.Equal(t, 0, len(needMigrateFiles))
}

func TestGetNeedRollbackKeys(t *testing.T) {
	err := db.Def().Exec("truncate migrate_file").Error
	assert.Nil(t, err)
	var needRollbackMs []migrate.MigrateFile
	needRollbackMs = migrate.GetNeedRollbackKeys(1)
	assert.Equal(t, 0, len(needRollbackMs))

	m := migrate.Migration{
		Migration: migrate.MigrateFiles[0].Key(),
		Batch:     1,
	}
	db.Def().Create(&m)

	needRollbackMs = migrate.GetNeedRollbackKeys(1)
	assert.Equal(t, 1, len(needRollbackMs))
	assert.Equal(t, migrate.MigrateFiles[0].Key(), needRollbackMs[0].Key())
}

func TestGetNextBatchNo(t *testing.T) {
	var nextBatch uint
	err := db.Def().Exec("truncate migrate_file").Error
	assert.Nil(t, err)
	nextBatch = migrate.GetNextBatchNo()
	assert.Equal(t, uint(1), nextBatch)
	m := migrate.Migration{
		Migration: migrate.MigrateFiles[0].Key(),
		Batch:     nextBatch,
	}
	db.Def().Create(&m)

	nextBatch = migrate.GetNextBatchNo()
	assert.Equal(t, uint(2), nextBatch)
}
