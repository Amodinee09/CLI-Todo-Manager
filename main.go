package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Amodinee09/todo-cli/todo"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: add | list | done | delete")
		return
	}

	manager := todo.Manager{}
	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task")
			return
		}

		task := os.Args[2]
		manager.AddTodo(task)
		manager.ListTodos()

	case "list":
		manager.ListTodos()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide todo ID")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])
		manager.MarkDone(id)
		manager.ListTodos()

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide todo ID")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])
		manager.DeleteTodo(id)
		manager.ListTodos()

	default:
		fmt.Println("Unknown command")
	}
}
