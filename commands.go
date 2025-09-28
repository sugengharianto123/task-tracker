package main

import (
	"errors"
	"fmt"
	"time"
)

var validStatuses = map[string]bool{
	"todo":        true,
	"in-progress": true,
	"done":        true,
}

// addTask menambahkan tugas baru dengan deskripsi yang diberikan
func AddTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	id := nextID(tasks)
	newTask := NewTask(id, description)
	tasks = append(tasks, newTask)
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("Task added successfully with ID %d\n", id)
	return nil
}

// listTasks mencetak semua tugas, atau yang difilter berdasarkan status
func ListTasks(filter string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	if filter != "" {
		if !validStatuses[filter] {
			fmt.Println("Invalid status filter. Valid statuses are: todo, in-progress, done")
			return
		}
		var out []Task
		for _, t := range tasks {
			if t.Status == filter {
				out = append(out, t)
			}
		}
		if len(out) == 0 {
			fmt.Println("No tasks found with status:", filter)
			return
		}
		for _, t := range out {
			printTask(t)
		}
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range tasks {
		printTask(t)
	}
}

func UpdateTask(id int, newDesc string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	i, err := findTaskIndex(tasks, id)
	if err != nil {
		return fmt.Errorf("Task with ID %d not found", id)
	}
	tasks[i].Description = newDesc
	tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("Task with ID %d updated successfully\n", id)
	return nil
}

func DeleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	i, err := findTaskIndex(tasks, id)
	if err != nil {
		return fmt.Errorf("Task with ID %d not found", id)
	}
	tasks = append(tasks[:i], tasks[i+1:]...)
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("Task with ID %d deleted successfully\n", id)
	return nil
}

func MarkTask(id int, status string) error {
	if !validStatuses[status] {
		return errors.New("Invalid status")
	}
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	i, err := findTaskIndex(tasks, id)
	if err != nil {
		return fmt.Errorf("Task with ID %d not found", id)
	}
	tasks[i].Status = status
	tasks[i].UpdatedAt = time.Now().Format(time.RFC3339)
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("Task with ID %d marked as %s successfully\n", id, status)
	return nil
}
