package task

import "fmt"

type TaskManager struct {
	Tasks []Task
}

func NewTaskManager(tasks []Task) TaskManager {
	return TaskManager{
		Tasks: tasks,
	}
}

func (tm *TaskManager) AddTask(name string) {
	id := len(tm.Tasks) + 1
	task := NewTask(id, name)
	tm.Tasks = append(tm.Tasks, task)
	fmt.Println("new task added -> ", name)
}

// List not completed tasks
func (tm *TaskManager) ListTask(showAll bool) {

	if showAll {
		for _, task := range tm.Tasks {
			fmt.Printf("%d\t%s\t%t\t%s\n", task.ID, task.Content, task.Completed, task.DateCreated)
		}
	} else {
		for _, task := range tm.Tasks {
			if task.Completed {
				continue
			}
			fmt.Printf("%d\t%s\t%t\t%s\n", task.ID, task.Content, task.Completed, task.DateCreated)
		}
	}
}

func (tm *TaskManager) ListAllTask() {
}

func (tm *TaskManager) RemoveTask(id int) {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			fmt.Println("task removed -> ", task.Content)
			return
		}
	}
	fmt.Println("task not found. id -> ", id)
}

func (tm *TaskManager) CompleteTask(id int) {
	for i, task := range tm.Tasks {
		if task.ID == id && !task.Completed {
			tm.Tasks[i].Completed = true
			return
		} else if task.ID == id && task.Completed {
			fmt.Println("task is already completed.")
		}
	}
}
