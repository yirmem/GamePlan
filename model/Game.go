package model

import (
	"gameplan/db"
	"log/slog"
)

type Game struct {
	Id   int64  `json:"id" xorm:"pk autoincr"`
	Name string `json:"name" xorm:"varchar(50) comment('游戏名称')"`
}

func (Game) TableName() string {
	return "game"
}

func (game Game) Insert() {
	_, err := db.SqliteDb.Insert(game)
	if err != nil {
		slog.Error("添加游戏错误", "", err)
	}
}

func (game Game) Update() {
	_, err := db.SqliteDb.Cols("name").ID(game.Id).Update(game)
	if err != nil {
		slog.Error("修改游戏错误", "", err)
	}
}

func (game Game) Delete() {
	//删除游戏
	_, err := db.SqliteDb.ID(game.Id).Delete(&game)
	if err != nil {
		slog.Error("删除游戏错误", "", err)
	}
	//删除游戏所有任务
	task := &Task{}
	_, err = db.SqliteDb.Where("game_id=?", game.Id).Delete(task)
	if err != nil {
		slog.Error("删除游戏任务列表错误", "", err)
	}
}
