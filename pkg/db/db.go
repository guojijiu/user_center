package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"user_center/config"
	"user_center/pkg/glog"
	_ "github.com/go-sql-driver/mysql"
)

var connections = map[string]*gorm.DB{}

func InitDef() (*gorm.DB, error) {
	db, err := InitConn(config.Database.MySQL["default"])
	if err != nil {
		return db, err
	}
	connections["default"] = db
	return db, nil
}

func InitConn(c config.MysqlConf) (*gorm.DB, error) {
	var conn *gorm.DB
	var err error
	conn, err = gorm.Open("mysql", c.String())
	if err != nil {
		return conn, err
	}
	conn.LogMode(config.Debug)
	conn.SetLogger(glog.Channel("db"))
	conn.DB().SetConnMaxLifetime(c.MaxLiftTime)
	return conn, nil
}

func Close() {
	for k, conn := range connections {
		if err := conn.Close(); err != nil {
			log.Printf("Close mysql conn %s err: %+v", k, err)
		}
	}
}

// get default conn
func Def() *gorm.DB {
	return Conn("default")
}

// 获取指定的连接, 当创建新连接失败时, 将会返回默认连接
func Conn(name string) *gorm.DB {
	var err error
	if conn, ok := connections[name]; ok {
		return conn
	}
	if c, ok := config.Database.MySQL[name]; ok {
		connections[name], err = InitConn(c)
		if err != nil {
			panic(fmt.Sprintf("Connect mysql (%s: %s) failed: %+v", name, c.String(), err))
			return nil
		}
		return connections[name]
	}
	panic(fmt.Sprintf("Can't read mysql config: %s", name))
	return nil
}
