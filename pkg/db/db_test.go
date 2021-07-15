package db_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"user_center/pkg/db"
)

func TestConn(t *testing.T) {
	var err error
	_, err = db.InitDef()
	assert.Nil(t, err)
	//err = db.Def().DB().Ping()
	assert.Nil(t, err)
}

func TestQuery(t *testing.T) {
	var err error
	_, err = db.InitDef()
	assert.Nil(t, err)
	//assert.Nil(t, db.Def().DB().Ping())
}

func TestGetTables(t *testing.T) {
	var err error
	tables := []string{}
	err = db.Def().Raw("show tables").Pluck("Tables_in_mysql", &tables).Error
	assert.Nil(t, err)
	t.Log(tables)
}
