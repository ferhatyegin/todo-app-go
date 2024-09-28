package storage

import (
	"encoding/json"
	"os"

	"github.com/ferhatyegin/todo-app-go/internal/task"
)

func LoadTasks(filepath string) ([]task.Task, error) {
	taskFile, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}
	defer taskFile.Close()

	var tasks []task.Task
	decoder := json.NewDecoder(taskFile)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(filepath string, tasks []task.Task) error {
	taskFile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer taskFile.Close()

	encoder := json.NewEncoder(taskFile)
	if err := encoder.Encode(tasks); err != nil {
		return err
	}
	return nil
}
