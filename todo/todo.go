package todo

import (
	"time"
)

type Todo struct {
	ID        int
	Task      string
	Status    string
	CreatedAt time.Time
}
