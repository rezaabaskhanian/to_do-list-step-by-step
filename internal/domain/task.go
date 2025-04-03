package domain

// مدل هایی مه در دنیای واقعی هستن را پیاده سازی میکنیم
// مثلا هر تسک در دنیای واقعی شامل موارد زیر است

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
