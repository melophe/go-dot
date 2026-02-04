package main

import (
	"fmt"
	"os"
)

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
