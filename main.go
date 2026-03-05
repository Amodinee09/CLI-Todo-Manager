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
	manager.Load() // load existing todos

	command := os.Args[1]

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task")
			return
		}

		task := os.Args[2]

		id := manager.AddTodo(task)
		manager.Save()

		fmt.Println("Todo added successfully with ID:", id)

	case "list":
		manager.ListTodos()

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide todo ID")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])

		ok := manager.MarkDone(id)

		if ok {
			manager.Save()
			fmt.Printf("✔ Todo %d marked as done\n", id)
		} else {
			fmt.Println("Todo not found")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide todo ID")
			return
		}

		id, _ := strconv.Atoi(os.Args[2])

		ok := manager.DeleteTodo(id)

		if ok {
			manager.Save()
			fmt.Printf("✔ Todo %d deleted successfully\n", id)
		} else {
			fmt.Println("Todo not found")
		}

	default:
		fmt.Println("Unknown command")
	}
}
