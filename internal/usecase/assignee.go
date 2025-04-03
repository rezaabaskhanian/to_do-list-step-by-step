package usecase

import "yourproject/internal/domain"

func CreateTask(title, description string) domain.Task {
	task := domain.Task{
		ID:          1, // در آینده باید این ID از دیتابیس یا حافظه گرفته بشه
		Title:       title,
		Description: description,
		Done:        false,
	}
	return task
}
