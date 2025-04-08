package usecase

import (
	"time"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
)

type TaskUseCase struct {
	repo domain.TaskRepository
}

// NewTaskUseCase یک نمونه جدید از TaskUseCase ایجاد می‌کند.
func NewTaskUseCase(repo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

// CreateTask یک تسک جدید ایجاد کرده و در ذخیره‌سازی می‌نویسد.
func (u *TaskUseCase) CreateTask(title, description string) (domain.Task, error) {
	tasks, err := u.repo.Load()
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

	if err := u.repo.Save(tasks); err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

// ListTasks لیست کامل تسک‌ها را برمی‌گرداند.
func (u *TaskUseCase) ListTasks() ([]domain.Task, error) {
	return u.repo.Load()
}
