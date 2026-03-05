package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const fileName = "todos.json"

type Manager struct {
	todos   []Todo
	counter int
}

func (m *Manager) Load() {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	json.Unmarshal(data, &m.todos)

	for _, t := range m.todos {
		if t.ID > m.counter {
			m.counter = t.ID
		}
	}
}

func (m *Manager) Save() {

	data, _ := json.MarshalIndent(m.todos, "", "  ")
	os.WriteFile(fileName, data, 0644)
}

func (m *Manager) AddTodo(task string) int {

	m.counter++
	id := m.counter

	t := Todo{
		ID:        id,
		Task:      task,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	m.todos = append(m.todos, t)

	return id
}

func (m *Manager) ListTodos() {

	if len(m.todos) == 0 {
		fmt.Println("No todos found")
		return
	}

	fmt.Printf("%-5s %-20s %-10s %-20s\n", "ID", "TASK", "STATUS", "CREATED")

	for _, v := range m.todos {
		fmt.Printf("%-5d %-20s %-10s %-20s\n",
			v.ID,
			v.Task,
			v.Status,
			v.CreatedAt.Format("2006-01-02"),
		)
	}
}
func (m *Manager) MarkDone(id int) bool {

	for i := range m.todos {
		if m.todos[i].ID == id {
			m.todos[i].Status = "done"
			return true
		}
	}

	return false
}
func (m *Manager) DeleteTodo(id int) bool {

	for i := range m.todos {
		if m.todos[i].ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return true
		}
	}

	return false
}
