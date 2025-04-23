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

