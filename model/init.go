package model

import (
	"gameplan/db"
	"log/slog"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func Init() {

	c := cron.New()

	//启动时立即执行一次
	resetGameTask(false)
	// 每天0点执行
	_, err := c.AddFunc("*/1 * * * *", func() {
		resetGameTask(true)
	})
	if err != nil {
		panic(err)
	}

	c.Start()
	db.SqliteDb.Sync2(&Game{}, &Task{})

}
func resetGameTask(event bool) {

	taskList := make([]Task, 0)
	err := db.SqliteDb.Where("is_repeat = 1").Find(&taskList)
	if err != nil {
		slog.Error("查询待重置任务失败", "error", err)
		return
	}
	if len(taskList) == 0 {
		return
	}

	now := time.Now()
	for i := range taskList {
		task := &taskList[i]

		// 解析当前周期起始时间
		startTime, err := time.ParseInLocation("2006-01-02 15:04:05", task.StartDate, time.Local)
		if err != nil {
			slog.Error("解析开始时间失败",
				"task_id", task.Id,
				"start_date", task.StartDate,
				"error", err)
			continue
		}

		cycleDays := task.CheckDay
		if cycleDays <= 0 {
			cycleDays = 1
		}
		cycleDuration := time.Duration(cycleDays) * 24 * time.Hour

		// 当前周期的结束时间点
		periodEnd := startTime.Add(cycleDuration)

		// 如果当前时间还未到达当前周期结束点，无需重置
		if now.Before(periodEnd) {
			continue
		}

		// 需要重置
		// 下一个周期开始时间 = startTime + (periodsPassed + 1) * cycleDuration
		periodsPassed := int64(now.Sub(startTime) / cycleDuration)
		nextStart := startTime
		for i := 0; i < int(periodsPassed); i++ {
			nextStart = nextStart.Add(cycleDuration)
		}

		// 更新内存中的任务
		task.Status = 0
		task.StartDate = nextStart.Format("2006-01-02 15:04:05")

		// 更新数据库（只更新 status 和 start_date）
		affected, err := db.SqliteDb.ID(task.Id).
			Cols("status", "start_date").
			Update(task)
		if err != nil {
			slog.Error("重置任务状态失败", "task_id", task.Id, "error", err)
			continue
		}
		if affected > 0 {
			slog.Info("重置任务成功",
				"task_id", task.Id,
				"name", task.Name,
				"old_start", startTime.Format("2006-01-02 15:04:05"),
				"new_start", task.StartDate)
		}
	}
	if event {
		app := application.Get()
		app.Event.Emit("task_reset")
	}
}
