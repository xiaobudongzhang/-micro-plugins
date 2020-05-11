package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/xiaobudongzhang/micro-basic/config"

	"github.com/micro/go-micro/v2/util/log"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("db 已经初始化")
		log.Logf(err.Error())
		return
	}

	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
