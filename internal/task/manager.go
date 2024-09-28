package task

import (
	"fmt"
	"os"
	"text/tabwriter"
)

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

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tContent\tCompleted\tDate Created")

	if showAll {
		for _, task := range tm.Tasks {
			fmt.Fprintf(w, "%-5d\t%-20s\t%-10t\t%s\n", task.ID, task.Content, task.Completed, task.DateCreated)
		}
		w.Flush()
	} else {
		for _, task := range tm.Tasks {
			if task.Completed {
				continue
			}
			fmt.Fprintf(w, "%-5d\t%-20s\t%-10t\t%s\n", task.ID, task.Content, task.Completed, task.DateCreated)
		}
		w.Flush()
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
