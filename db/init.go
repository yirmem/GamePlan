package db

import (
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var SqliteDb *xorm.Engine

func Init() {
	var err error
	SqliteDb, err = xorm.NewEngine("sqlite3", "./gp_data/gp.db")
	if err != nil {
		slog.Error("初始化sqlite失败", "", err.Error())
	}

}
