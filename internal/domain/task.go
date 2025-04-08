package domain

import "time"

// مدل هایی مه در دنیای واقعی هستن را پیاده سازی میکنیم
// مثلا هر تسک در دنیای واقعی شامل موارد زیر است

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Done        bool      `json:"done"`
}

// TaskRepository مسئول ذخیره و بازیابی تسک‌ها از یک منبع داده است.
type TaskRepository interface {
	// Save همه‌ی تسک‌ها را ذخیره می‌کند (مثلاً در یک فایل یا دیتابیس)
	Save(tasks []Task) error

	// Load لیست تسک‌ها را از منبع داده بارگذاری می‌کند
	Load() ([]Task, error)
}
