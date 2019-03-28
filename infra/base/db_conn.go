package base

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

// 基础资源访问,算法 建议使用公共函数

var (
	database sqlbuilder.Database
)

func UpperDatabase() sqlbuilder.Database {
	if database == nil {
		InitUpperDatabase()
	}
	return database
}

//database定义为interface{} 需要被实例化
func InitUpperDatabase() {
	settings := mysql.ConnectionURL{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "root",
		Database: "redenvelope",
	}
	db, err := mysql.Open(settings)
	if err != nil {
		panic(err)
	}
	database = db
}
