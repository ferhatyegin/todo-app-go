package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ferhatyegin/todo-app-go/internal/storage"
	"github.com/ferhatyegin/todo-app-go/internal/task"
)

func main() {

	//If there were no flags given print out manual and shutdown
	if len(os.Args) < 2 {
		fmt.Println(`
Add: Adds a new task to the end of list.
	add <string:content>

List: List all tasks.
	list

Remove: Removes a task by id [default value: -1]
	remove <int:id>

Complete: Completes a task by id
	complete <int:id>
		`)
		os.Exit(1)
	}

	tasks, err := storage.LoadTasks("tasks.json")
	if err != nil {
		fmt.Println("error while loading tasks: ", err)
		os.Exit(1)
	}

	taskManager := task.NewTaskManager(tasks)

	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println(`Add: Adds a new task to the end of list.
	add <string>`)
			os.Exit(1)
		}
		taskManager.AddTask(os.Args[2])

	case "list":
		showAll := false
		if len(os.Args) > 2 && os.Args[2] == "-a"{
			showAll = true
		}
		taskManager.ListTask(showAll)

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println(`Remove: Removes a task by id [default value: -1]
	remove <int>`)
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("invalid task ID")
			os.Exit(1)
		}

		taskManager.RemoveTask(id)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println(``)
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid task ID")
		}
		taskManager.CompleteTask(id)

	default:
		fmt.Println(`Unkown Command
		
Add: Adds a new task to the end of list.
	add <string>

List: List all tasks.
	list

Remove: Removes a task by id [default value: -1]
	remove <int>
		`)
		os.Exit(1)
	}

	err = storage.SaveTasks("tasks.json", taskManager.Tasks)
	if err != nil {
		fmt.Println("error while saving tasks: ", err)
		os.Exit(1)
	}

}
