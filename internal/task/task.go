package task

import "time"

type Task struct {
	ID          int    `json:"id"`
	Content     string `json:"content"`
	Completed   bool   `json:"completed"`
	DateCreated string `json:"dateCreated"`
}

func NewTask(id int, content string) Task {
	return Task{
		ID:          id,
		Content:     content,
		Completed:   false,
		DateCreated: time.Now().Format("02-01-2006 15:04"),
	}
}
