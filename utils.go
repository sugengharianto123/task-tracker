package main

import (
	"fmt"
)

// printTask mencetak detail tugas
func printTask(t Task){
	fmt.Printf("ID : %d", t.ID)
	fmt.Printf("\nDescription : %s", t.Description)
	fmt.Printf("\nStatus : %s", t.Status)
	fmt.Printf("\nCreated At : %s", t.CreatedAt)
	fmt.Printf("\nUpdated At : %s\n", t.UpdatedAt)
	fmt.Println("===================================")
}

// printUsage mencetak penggunaan program
func printUsage(){
	fmt.Println("Task Tracker CLI (GO)")
	fmt.Println("Usage:")
	fmt.Println(" task-cli add \"Buy groceries\"")
	fmt.Println(" task-cli list")
	fmt.Println(" task-cli list todo|in-progress|done")
	fmt.Println(" task-cli update <id> \"new description\"")
	fmt.Println(" task-cli delete <id>")
	fmt.Println(" task-cli mark-in-progress <id>")
	fmt.Println(" task-cli mark-done <id>")
}