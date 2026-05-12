package model

import (
	"gameplan/db"
	"log/slog"
)

// 游戏任务
type Task struct {
	Id        int64  `json:"id" xorm:"pk autoincr"`
	Name      string `json:"name" xorm:"varchar(200) comment('任务名称')"`
	CheckDay  int    `json:"checkDay" xorm:"comment('任务持续时间 单位是天')"`
	IsRepeat  int    `json:"isRepeat" xorm:"comment('是否重复 单次0/重复1')"`
	StartDate string `json:"startDate" xorm:"varchar(50) comment('任务初始生效日期')"`
	GameId    int64  `json:"gameId" xorm:"comment('游戏Id')"`
	Content   string `json:"content" xorm:"varchar(1000) comment('自定义任务内容或奖励')"`
	Status    int    `json:"status" xorm:"comment('完成状态 1完成/0未完成')"`
	Updated   string `json:"updated" xorm:"updated varchar(50) comment('最后更新时间')"`
	Created   string `json:"created" xorm:"created varchar(50) comment('创建时间')"`
}

func (Task) TableName() string {
	return "game_task"
}

func (task Task) GetList() ([]Task, int64) {
	result := make([]Task, 0)
	sql := db.SqliteDb.Select("*")
	if task.CheckDay != 0 {
		sql.And("check_day=?", task.CheckDay)
	}
	if task.GameId != 0 {
		sql.And("game_id=?", task.GameId)
	}
	if task.Status == -1 {
		sql.And("status=0")
	}

	count, err := sql.FindAndCount(&result)
	if err != nil {
		slog.Error("查询游戏任务列表错误", "", err)
	}
	return result, count
}

func (task Task) Insert() {
	_, err := db.SqliteDb.Insert(task)
	if err != nil {
		slog.Error("添加游戏任务错误", "", err)
	}
}

func (task Task) Update() {
	sql := db.SqliteDb.ID(task.Id)
	_, err := sql.MustCols("status,is_repeat").Update(&task)
	if err != nil {
		slog.Error("修改游戏任务错误", "", err)
	}
}

func (task Task) Delete() {
	_, err := db.SqliteDb.ID(task.Id).Delete(&task)
	if err != nil {
		slog.Error("删除游戏错误", "", err)
	}
}
