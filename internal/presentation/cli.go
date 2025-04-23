package cli

//struct cli

// NEW

// function haro minvisi

//inja dige to faghat fucntion haye laye service ro az interface khodet call mikoni

import (
	"fmt"

	Assignee "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase/assignee"
	Task "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase/task"
)

// تعریف دستورات CLI
type CLI struct {
	taskService     *Task.TaskService
	assigneeService *Assignee.AssigneeService
}

// تابع تنظیم CLI برای اتصال سرویس‌ها
func SetupCLI(taskService *Task.TaskService, assigneeService *Assignee.AssigneeService) {
	cli := CLI{
		taskService:     taskService,
		assigneeService: assigneeService,
	}

	// نمایش دستوراتی که کاربر می‌تواند انجام دهد
	cli.showMenu()

	// بعد از این باید از ورودی کاربر بخواهیم و دستورات را اجرا کنیم
}

// تابعی برای نمایش منو به کاربر
func (cli *CLI) showMenu() {
	fmt.Println("لطفاً یکی از دستورات زیر را وارد کنید:")
	fmt.Println("1. ایجاد تسک")
	fmt.Println("2. نمایش لیست تسک‌ها")
	fmt.Println("3. ایجاد مسئول")
	fmt.Println("4. نمایش لیست مسئولین")
	fmt.Println("5. تغییر وضعیت تسک به انجام‌شده")
	fmt.Println("6. حذف تسک")
	fmt.Println("7. خروج")

	var choice int
	fmt.Scanln(&choice)

	// اجرا دستورات انتخابی کاربر
	cli.executeCommand(choice)
}

// تابع برای اجرای دستورات انتخابی کاربر
func (cli *CLI) executeCommand(choice int) {
	switch choice {
	case 1:
		cli.createTask()
	case 2:
		cli.listTasks()
	case 3:
		cli.createAssignee()
	case 4:
		cli.listAssignees()
	case 5:
		cli.markTaskAsDone()
	case 6:
		cli.deleteTask()
	case 7:
		fmt.Println("خروج از برنامه")
		return
	default:
		fmt.Println("دستور نامعتبر. لطفاً دوباره تلاش کنید.")
		cli.showMenu()
	}
}

// تابع برای ایجاد تسک جدید
func (cli *CLI) createTask() {
	var title, description string
	fmt.Println("عنوان تسک را وارد کنید:")
	fmt.Scanln(&title)
	fmt.Println("توضیحات تسک را وارد کنید:")
	fmt.Scanln(&description)

	task, err := cli.taskService.CreateTask(title, description)
	if err != nil {
		fmt.Println("خطا در ایجاد تسک:", err)
		return
	}

	fmt.Printf("تسک با موفقیت ایجاد شد: %v\n", task)
	cli.showMenu()
}

// تابع برای نمایش لیست تسک‌ها
func (cli *CLI) listTasks() {
	tasks, err := cli.taskService.ListTasks()
	if err != nil {
		fmt.Println("خطا در بارگذاری تسک‌ها:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("هیچ تسکی وجود ندارد.")
	} else {
		fmt.Println("لیست تسک‌ها:")
		for _, t := range tasks {
			fmt.Printf("ID: %d, عنوان: %s, وضعیت: %v\n", t.ID, t.Title, t.Done)
		}
	}

	cli.showMenu()
}

// تابع برای ایجاد مسئول جدید
func (cli *CLI) createAssignee() {
	var name, email string
	fmt.Println("نام مسئول را وارد کنید:")
	fmt.Scanln(&name)
	fmt.Println("ایمیل مسئول را وارد کنید:")
	fmt.Scanln(&email)

	assignee, err := cli.assigneeService.CreateAssignee(name, email)
	if err != nil {
		fmt.Println("خطا در ایجاد مسئول:", err)
		return
	}

	fmt.Printf("مسئول با موفقیت ایجاد شد: %v\n", assignee)
	cli.showMenu()
}

// تابع برای نمایش لیست مسئولین
func (cli *CLI) listAssignees() {
	assignees, err := cli.assigneeService.ListAssignees()
	if err != nil {
		fmt.Println("خطا در بارگذاری مسئولین:", err)
		return
	}

	if len(assignees) == 0 {
		fmt.Println("هیچ مسئولیتی وجود ندارد.")
	} else {
		fmt.Println("لیست مسئولین:")
		for _, a := range assignees {
			fmt.Printf("ID: %d, نام: %s, ایمیل: %s\n", a.ID, a.Name, a.Email)
		}
	}

	cli.showMenu()
}

// تابع برای تغییر وضعیت تسک به انجام‌شده
func (cli *CLI) markTaskAsDone() {
	var taskID int
	fmt.Println("ID تسک را وارد کنید تا وضعیت آن تغییر کند:")
	fmt.Scanln(&taskID)

	err := cli.taskService.MarkTaskAsDone(taskID)
	if err != nil {
		fmt.Println("خطا در تغییر وضعیت تسک:", err)
		return
	}

	fmt.Println("وضعیت تسک با موفقیت به انجام‌شده تغییر یافت.")
	cli.showMenu()
}

// تابع برای حذف تسک
func (cli *CLI) deleteTask() {
	var taskID int
	fmt.Println("ID تسک را وارد کنید تا حذف شود:")
	fmt.Scanln(&taskID)

	err := cli.taskService.DeleteTask(taskID)
	if err != nil {
		fmt.Println("خطا در حذف تسک:", err)
		return
	}

	fmt.Println("تسک با موفقیت حذف شد.")
	cli.showMenu()
}

// تابع برای اجرای دستورات از ورودی کاربر
func Execute() {
	// در اینجا می‌توانید برنامه را برای دریافت ورودی و اجرای دستورات شروع کنید
	// برای مثال، از یک حلقه ورودی برای اجرا استفاده کنید.
}
