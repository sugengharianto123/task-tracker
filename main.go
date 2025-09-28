package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add \"description\"")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		if err := AddTask(description); err != nil {
			fmt.Println("Error adding task:", err)
		}
	case "list":
		if len(os.Args) == 3 {
			ListTasks(os.Args[2])
		} else {
			ListTasks("")
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> \"new description\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID:", os.Args[2])
			return
		}
		desc := strings.Join(os.Args[3:], " ")
		if err := UpdateTask(id, desc); err != nil {
			fmt.Println("Error updating task:", err)
		}
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id is must be a number")
			return
		}
		if err := DeleteTask(id); err != nil {
			fmt.Println("Error deleting task:", err)
		}
	case "mark-in-progress":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id is must be a number")
			return
		}
		if err := MarkTask(id, "in-progress"); err != nil {
			fmt.Println("Error marking task as in-progress:", err)
		}
	case "mark-done":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id is must be a number")
			return
		}
		if err := MarkTask(id, "done"); err != nil {
			fmt.Println("Error marking task as done:", err)
		}
	default:
		fmt.Println("Unknown command:", cmd)
		printUsage()
	}
}
