package domain

// مدل هایی مه در دنیای واقعی هستن را پیاده سازی میکنیم
// مثلا هر تسک در دنیای واقعی شامل موارد زیر است

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool // مشخص می‌کنه که تسک انجام شده یا نه
}
