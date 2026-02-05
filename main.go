package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Task struct {
	ID	  int
	Name string
	Done  bool
}

const taskFile = "tasks.json"

// Save tasks to JSON file
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

// Load tasks from JSON file
func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func main() {
	// Check args: need at least a command
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-do <command> [args]")
		return
	}

	// Command handling
	switch os.Args[1] {
	case "add":
		fmt.Println("Adding task...")
	case "list":
		fmt.Println("Listing tasks...")
	case "done":
		fmt.Println("Completing task...")
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
