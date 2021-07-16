package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user_center/config"
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
	// 初始化sql日志log
	initSqlLog()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.String(), // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
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
