package Task

import (
	"fmt"
	"time"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
)

type TaskRepositoryFile struct{}

func (r *TaskRepositoryFile) Save(tasks []domain.Task) error {
	return infrastructure.SaveTasks(tasks)
}

func (r *TaskRepositoryFile) Load() ([]domain.Task, error) {
	return infrastructure.LoadTasks()
}

type TaskService struct {
	taskRepository TaskRepository
}

// type TaskUseCase struct {
// 	repo TaskRepository
// }

func NewTaskService(taskRepository TaskRepository) *TaskService {
	// return &TaskService{taskRepository: TaskRepository}
	return &TaskService{taskRepository: taskRepository}
}

//   - user can create a task and assigne
//
// CreateTask یک تسک جدید ایجاد کرده و در ذخیره‌سازی می‌نویسد.
func (u *TaskService) CreateTask(title, description string) (domain.Task, error) {
	tasks, err := u.taskRepository.Load()
	if err != nil {
		return domain.Task{}, err
	}

	// یافتن بیشترین ID برای تولید ID جدید
	newID := 1
	for _, t := range tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}

	task := domain.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	tasks = append(tasks, task)

	if err := u.taskRepository.Save(tasks); err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

//TODO://     // - user can change the task status to done

//TODO://	   //- user can delete a task

//TODO:  // user can list task

// ListTasks لیست کامل تسک‌ها را برمی‌گرداند.
func (u *TaskService) ListTasks() ([]domain.Task, error) {
	return u.taskRepository.Load()
}

// MarkTaskAsDone یک تسک را به حالت انجام‌شده تغییر می‌دهد.
func (u *TaskService) MarkTaskAsDone(taskID int) error {
	tasks, err := u.taskRepository.Load()
	if err != nil {
		return err
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("تسک با ID %d یافت نشد", taskID)
	}

	return u.taskRepository.Save(tasks)
}

// DeleteTask یک تسک را از لیست حذف می‌کند.
func (u *TaskService) DeleteTask(taskID int) error {
	tasks, err := u.taskRepository.Load()
	if err != nil {
		return err
	}

	newTasks := make([]domain.Task, 0, len(tasks))
	found := false
	for _, t := range tasks {
		if t.ID == taskID {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return fmt.Errorf("تسک با ID %d یافت نشد", taskID)
	}

	return u.taskRepository.Save(newTasks)
}
