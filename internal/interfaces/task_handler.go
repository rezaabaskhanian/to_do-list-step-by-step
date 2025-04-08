package interfaces

import (
	"flag"
	"fmt"
	"os"

	// "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase"
)

type TaskHandler struct {
	usecase *usecase.TaskUseCase
}

func NewTaskHandler(uc *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{usecase: uc}
}

func (h *TaskHandler) CreateTask(title, description string) {
	task, err := h.usecase.CreateTask(title, description)
	if err != nil {
		fmt.Println("خطا در ایجاد تسک:", err)
		os.Exit(1)
	}
	fmt.Printf("✅ تسک ایجاد شد: %+v\n", task)
}

func (h *TaskHandler) ListTasks() {
	tasks, err := h.usecase.ListTasks()
	if err != nil {
		fmt.Println("خطا در بارگذاری تسک‌ها:", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("❕ هیچ تسکی موجود نیست.")
		return
	}

	for _, task := range tasks {
		fmt.Printf("📝 ID: %d | عنوان: %s | توضیحات: %s | وضعیت: %v\n",
			task.ID, task.Title, task.Description, task.Done)
	}
}

func RunCLI(usecase *usecase.TaskUseCase) {
	taskHandler := NewTaskHandler(usecase)

	createTaskCmd := flag.NewFlagSet("create-task", flag.ExitOnError)
	title := createTaskCmd.String("title", "", "عنوان تسک")
	description := createTaskCmd.String("description", "", "توضیحات تسک")

	listTaskCmd := flag.NewFlagSet("list-tasks", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("لطفاً یک دستور وارد کنید: create-task یا list-tasks")
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
		fmt.Println("دستور نامعتبر است. از create-task یا list-tasks استفاده کنید.")
	}
}
