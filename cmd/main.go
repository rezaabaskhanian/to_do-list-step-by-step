package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	// "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("یک ساب‌کامند مورد نیاز است: create-task یا list-tasks")
		os.Exit(1)
	}

	switch os.Args[1] {

	// go run main.go create-task -title="عنوان تسک" -description="توضیحات تسک"

	case "create-task":
		createTaskCmd := flag.NewFlagSet("create-task", flag.ExitOnError)
		title := createTaskCmd.String("title", "", "عنوان تسک")
		description := createTaskCmd.String("description", "", "توضیحات تسک")

		createTaskCmd.Parse(os.Args[2:])

		if *title == "" || *description == "" {
			fmt.Println("عنوان و توضیحات تسک نمی‌توانند خالی باشند.")
			createTaskCmd.PrintDefaults()
			os.Exit(1)
		}

		task := usecase.CreateTask(*title, *description)
		fmt.Printf("تسک ایجاد شد: %+v\n", task)

		// ذخیره تسک در فایل
		tasks, err := infrastructure.LoadTasks()
		if err != nil {
			fmt.Println("خطا در بارگذاری تسک‌ها:", err)
			os.Exit(1)
		}

		tasks = append(tasks, task)
		if err := infrastructure.SaveTasks(tasks); err != nil {
			fmt.Println("خطا در ذخیره تسک‌ها:", err)
			os.Exit(1)
		}

		//go run main.go list-tasks

	case "list-tasks":
		listTasksCmd := flag.NewFlagSet("list-tasks", flag.ExitOnError)
		listTasksCmd.Parse(os.Args[2:])

		tasks, err := infrastructure.LoadTasks()
		if err != nil {
			fmt.Println("خطا در بارگذاری تسک‌ها:", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("هیچ تسکی وجود ندارد.")
			return
		}

		fmt.Println("لیست تسک‌ها:")
		for _, task := range tasks {
			taskJSON, _ := json.MarshalIndent(task, "", "  ")
			fmt.Println(string(taskJSON))
		}

	default:
		fmt.Println("ساب‌کامند نامعتبر است. از create-task یا list-tasks استفاده کنید.")
		os.Exit(1)
	}
}
