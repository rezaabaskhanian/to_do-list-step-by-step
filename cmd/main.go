package main

import (
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must provide a command")
		return
	}

	command := os.Args[1]

	if command == "create-task" {
		if len(os.Args) < 4 {
			fmt.Println("Please provide a title and description for the task.")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		task := usecase.CreateTask(title, description)
		fmt.Printf("Task Created: %+v\n", task)
	}
}
