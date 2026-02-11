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
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-do add <task name>")
			return
		}
		tasks, _ := loadTasks()
		newTask := Task{
			ID:	  len(tasks) + 1,
			Name: os.Args[2],
			Done: false,
		}
		tasks = append(tasks, newTask)
		saveTasks(tasks)
		fmt.Println("✓ Added:", newTask.Name)
	case "list":
		tasks, _ := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks yet")
			return
		}
		for _, task := range tasks {
			status := "[ ]"
			if task.Done {
				status = "[x]"
			}
			fmt.Printf("%d. %s %s\n", task.ID, status, task.Name)
		} 
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-do done <task id>")
			return
		}
		id := 0
		fmt.Sscanf(os.Args[2], "%d", &id)
		
		tasks, _ := loadTasks()
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Done = true
				saveTasks(tasks)
				fmt.Println("✓ Completed:", task.Name)
				return
			}
		}
		fmt.Println("Task not found:", id)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
