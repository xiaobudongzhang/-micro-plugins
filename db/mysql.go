package db

import (
	"database/sql"

	"github.com/xiaobudongzhang/micro-basic/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/util/log"
)

func initMysql() {
	var err error

	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
