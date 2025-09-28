package main

import "time"

// task adalah struktur data yang merepresentasikan sebuah tugas
// Task is a data structure that represents a task
type Task struct{
	ID		  int `json:"id"`
	Description     string `json:"description"`
	Status string `json:"status"` // todo, in-progress, done
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// NewTask membuat tugas baru dengan deskripsi yang diberikan
// NewTask Creates a new task with the given description and timestamps
func NewTask(id int, description string) Task {
	now := time.Now().Format(time.RFC3339)
	return Task {
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}