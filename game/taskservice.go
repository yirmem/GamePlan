package game

import (
	"encoding/json"
	"gameplan/db"
	"gameplan/model"
	"log/slog"
	"os"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type TaskService struct{}

func (g *TaskService) GetTaskList(task model.Task) ([]model.Task, int64) {
	return task.GetList()
}

func (g *TaskService) AddTask(task model.Task) {
	task.Insert()
}

func (g *TaskService) DelTask(task model.Task) {
	task.Delete()
}

func (g *TaskService) UpdateTask(task model.Task) {
	task.Update()
}

// 分享任务列表
func (g *TaskService) ShareTask(task model.Task) {
	game := &model.Game{}
	game.Id = task.GameId
	db.SqliteDb.Select("*").Where("id=?", task.GameId).Get(game)

	taskList := make([]model.Task, 0)
	db.SqliteDb.Select("name,check_day,is_repeat,start_date,content").Where("game_id=?", task.GameId).FindAndCount(&taskList)
	taskJson, _ := json.Marshal(taskList)

	app := application.Get()
	path, err := app.Dialog.SaveFile().
		AddFilter("JSON文件 (*.json)", "*.json").
		SetFilename(game.Name + ".json").
		PromptForSingleSelection()

	if path == "" {
		return // 用户取消了操作
	}

	// 3. 写入文件
	err = os.WriteFile(path, []byte(taskJson), 0644)
	if err != nil {
		return
	}

}

// 导入任务列表
func (g *TaskService) ImportTask(task model.Task) {
	app := application.Get()
	paths, err := app.Dialog.OpenFile().
		AddFilter("JSON文件 (*.json)", "*.json").
		PromptForMultipleSelection()

	if err != nil {
		return
	}

	for _, path := range paths {
		content, err := os.ReadFile(path)
		taskList := make([]model.Task, 0)
		json.Unmarshal(content, &taskList)
		for i := range taskList {
			singleTask := taskList[i]
			singleTask.GameId = task.GameId
			singleTask.Created = time.Now().Format(time.DateTime)
			singleTask.Insert()
		}
		if err != nil {
			slog.Error("导入文件错误", "", err.Error())
		}
	}
}
