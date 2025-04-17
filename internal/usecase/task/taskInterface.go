package Task

import "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"


// TaskRepository مسئول ذخیره و بازیابی تسک‌ها از یک منبع داده است.
type TaskRepository interface {
	// Save همه‌ی تسک‌ها را ذخیره می‌کند (مثلاً در یک فایل یا دیتابیس)
	Save(tasks []domain.Task) error

	// Load لیست تسک‌ها را از منبع داده بارگذاری می‌کند
	Load() ([]domain.Task, error)
}
