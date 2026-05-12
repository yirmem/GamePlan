package game

import (
	"gameplan/db"
	"gameplan/model"
	"log/slog"
)

type GameService struct{}

func (g *GameService) GetGameList() []model.Game {
	gameList := make([]model.Game, 0)
	_, err := db.SqliteDb.FindAndCount(&gameList)
	if err != nil {
		slog.Error("获取游戏列表错误", "", err)
	}
	return gameList
}

func (g *GameService) AddGame(game model.Game) {
	game.Insert()
}

func (g *GameService) DelGame(game model.Game) {
	game.Delete()
}

func (g *GameService) UpdateGame(game model.Game) {
	game.Update()
}
