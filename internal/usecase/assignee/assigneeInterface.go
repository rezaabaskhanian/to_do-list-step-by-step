package Assignee
import "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"


// AssigneeRepository مسئول ذخیره و بازیابی تسک‌ها از یک منبع داده است.
type AssigneeRepository interface {
	// Save همه‌ی تسک‌ها را ذخیره می‌کند (مثلاً در یک فایل یا دیتابیس)
	Save(tasks []domain.Assignee) error

	// Load لیست تسک‌ها را از منبع داده بارگذاری می‌کند
	Load() ([]domain.Assignee, error)
}
