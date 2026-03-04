package todo

import (
	"fmt"
)

type Manager struct {
	todos   []Todo
	counter int
}

func (m *Manager) AddTodo(task string) {
	m.counter++
	id := m.counter

	t := Todo{
		ID:     id,
		Task:   task,
		Status: "pending",
	}

	m.todos = append(m.todos, t)
}

func (m *Manager) ListTodos() {

	if len(m.todos) == 0 {
		fmt.Println("No todos found")
		return
	} else {
		for _, v := range m.todos {
			fmt.Println(v.ID, v.Task, v.Status, v.CreatedAt)
		}
	}
}

func (m *Manager) MarkDone(id int) {
	for i := range m.todos {
		if m.todos[i].ID == id {
			m.todos[i].Status = "done"
		}
	}
}

func (m *Manager) DeleteTodo(id int) {
	for i := range m.todos {
		if m.todos[i].ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			break
		}
	}
}
