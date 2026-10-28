package tasks

import (
	"errors"
	"time"
)

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidTitle = errors.New("title must be 1 to 200 characters")
)

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
