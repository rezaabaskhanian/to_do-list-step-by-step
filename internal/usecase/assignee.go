package usecase

import "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"

// CreateAssignee یک Assignee جدید ایجاد می‌کند
func CreateAssignee(name, email string) domain.Assignee {
	// به طور موقت ID را به صورت دستی می‌دهیم
	assignee := domain.Assignee{
		ID:    1, // در آینده باید این ID از دیتابیس یا حافظه گرفته بشه
		Name:  name,
		Email: email,
	}
	return assignee
}
