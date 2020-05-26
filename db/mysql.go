package db

import (
	"database/sql"

	"github.com/xiaobudongzhang/micro-basic/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2/util/log"
)

type db struct {
	Mysql Mysql `json:"mysql"`
}

// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func initMysql() {

	log.Logf("初始化mysql")

	c := config.C()
	cfg := &db{}

	err := c.App("db", cfg)

	if err != nil {
		log.Logf("initmysql err %s", err)
	}

	if !cfg.Mysql.Enable {
		log.Logf("initmysql 未启用 mysql")
		return
	}

	mysqlDB, err = sql.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	mysqlDB.SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)
	mysqlDB.SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
