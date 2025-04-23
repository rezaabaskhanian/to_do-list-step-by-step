package cli

import (
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	Assignee "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase/assignee"
	Task "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase/task"
)

// TaskInterfaceCli رابط برای تعاملات خط فرمان تسک‌ها
type TaskInterfaceCli interface {
	CreateTask(title, description string) (domain.Task, error)
	ListTasks() ([]domain.Task, error)
	UpdateTaskStatus(taskID int, done bool) error
	DeleteTask(taskID int) error
}

// AssigneeInterfaceCli رابط برای تعاملات خط فرمان مسئولین
type AssigneeInterfaceCli interface {
	CreateAssignee(name, email string) (domain.Assignee, error)
	ListAssignees() ([]domain.Assignee, error)
}

// TaskCli پیاده‌سازی TaskInterfaceCli برای تعاملات CLI
type TaskCli struct {
	taskService Task.TaskService
}

// NewTaskCli سازنده TaskCli
func NewTaskCli(taskService Task.TaskService) *TaskCli {
	return &TaskCli{taskService: taskService}
}

// CreateTask ایجاد تسک جدید
func (t *TaskCli) CreateTask(title, description string) (domain.Task, error) {
	return t.taskService.CreateTask(title, description)
}

// ListTasks لیست تمام تسک‌ها
func (t *TaskCli) ListTasks() ([]domain.Task, error) {
	return t.taskService.ListTasks()
}

// UpdateTaskStatus بروزرسانی وضعیت تسک (برای مثال انجام شده یا نشده)
// func (t *TaskCli) UpdateTaskStatus(taskID int, done bool) error {
// 	return t.taskService.UpdateTaskStatus(taskID, done)
// }

// DeleteTask حذف تسک
func (t *TaskCli) DeleteTask(taskID int) error {
	return t.taskService.DeleteTask(taskID)
}

// AssigneeCli پیاده‌سازی AssigneeInterfaceCli برای تعاملات CLI
type AssigneeCli struct {
	assigneeService Assignee.AssigneeService
}

// NewAssigneeCli سازنده AssigneeCli
func NewAssigneeCli(assigneeService Assignee.AssigneeService) *AssigneeCli {
	return &AssigneeCli{assigneeService: assigneeService}
}

// CreateAssignee ایجاد مسئول جدید
func (a *AssigneeCli) CreateAssignee(name, email string) (domain.Assignee, error) {
	return a.assigneeService.CreateAssignee(name, email)
}

// ListAssignees لیست تمام مسئولین
func (a *AssigneeCli) ListAssignees() ([]domain.Assignee, error) {
	return a.assigneeService.ListAssignees()
}
