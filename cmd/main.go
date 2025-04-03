package main

import (
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/interfaces"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("یک ساب‌کامند مورد نیاز است: create-task یا list-tasks")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-task":
		interfaces.CreateTask()
	case "list-tasks":
		interfaces.ListTasks()
	default:
		fmt.Println("ساب‌کامند نامعتبر است. از create-task یا list-tasks استفاده کنید.")
		os.Exit(1)
	}
}
