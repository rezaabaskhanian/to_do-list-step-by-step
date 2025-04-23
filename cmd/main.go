package main

import (
	"fmt"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure" // پکیج ذخیره‌سازی
	cli "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/presentation"
	Assignee "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase/assignee" // پکیج مسئولین
	Task "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase/task"         // پکیج تسک
)

func main() {

	// aval storage ro besaz ( init stroagte mikonam )

	// dovom boro application ro setup kon:
	//dar in marhale ma mikhayym ye struct repository todo ba ye assignee biarim bala
	//ke betonim be onvane vorudi bedim be taskService va assgineeService

	//hala barname kar mikone vali niaz dari ke behesh az ye tarifi vasl beshi.
	//mirim soraghe laye presentation va inja cli ro call mikonim

	// مرحله 1: راه‌اندازی ذخیره‌سازی
	err := infrastructure.InitializeStorage() // ایجاد فایل ذخیره‌سازی اگر موجود نباشد
	if err != nil {
		fmt.Println("خطا در راه‌اندازی ذخیره‌سازی:", err)
		return
	}

	// مرحله 2: راه‌اندازی رپوزیتوری‌ها و سرویس‌ها
	taskRepo := &Task.TaskRepositoryFile{}             // ساخت رپوزیتوری تسک
	assigneeRepo := &Assignee.AssigneeRepositoryFile{} // ساخت رپوزیتوری مسئولین

	taskService := Task.NewTaskService(taskRepo)                 // ایجاد سرویس تسک
	assigneeService := Assignee.NewAssigneeService(assigneeRepo) // ایجاد سرویس مسئولین

	// مرحله 3: اتصال سرویس‌ها به یکدیگر در لایه‌ی presentation
	cli.SetupCLI(taskService, assigneeService)

	// مرحله 4: اجرای برنامه از طریق CLI
	fmt.Println("در حال اجرای CLI...")
	cli.Execute()
}
