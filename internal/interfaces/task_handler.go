package interfaces

import (
	"flag"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase"
)

type TaskHandler struct {
	tasks []domain.Task
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (t *TaskHandler) CreateTask(title, description string) {

	task := usecase.CreateTask(title, description)
	fmt.Printf("تسک ایجاد شد: %+v\n", task)

	t.tasks = append(t.tasks, task)

	if err := infrastructure.SaveTasks(t.tasks); err != nil {
		fmt.Println("خطا در ذخیره تسک‌ها:", err)
		os.Exit(1)
	}
}

func (t *TaskHandler) ListTasks() {

	tasks, err := infrastructure.LoadTasks()
	if err != nil {
		fmt.Println("خطا در بارگذاری تسک‌ها:", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("هیچ تسکی موجود نیست.")
		return
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d, عنوان: %s, توضیحات: %s, وضعیت: %v\n", task.ID, task.Title, task.Description, task.Done)
	}
}

func RunCLI() {
	// ایجاد نمونه از TaskHandler
	taskHandler := NewTaskHandler()

	createTaskCmd := flag.NewFlagSet("create-task", flag.ExitOnError)
	title := createTaskCmd.String("title", "", "عنوان تسک")
	description := createTaskCmd.String("description", "", "توضیحات تسک")

	listTaskCmd := flag.NewFlagSet("list-tasks", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("لطفاً دستور را وارد کنید.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-task":
		createTaskCmd.Parse(os.Args[2:])
		if *title == "" || *description == "" {
			fmt.Println("عنوان و توضیحات تسک نمی‌توانند خالی باشند.")
			createTaskCmd.PrintDefaults()
			os.Exit(1)
		}

		taskHandler.CreateTask(*title, *description)

	case "list-tasks":
		listTaskCmd.Parse(os.Args[2:])

		taskHandler.ListTasks()

	default:
		fmt.Println("دستور نامعتبر است.")
		os.Exit(1)
	}
}
