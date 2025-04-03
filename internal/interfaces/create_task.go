package interfaces

import (
	"flag"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase"
)

func CreateTask() {
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
}
