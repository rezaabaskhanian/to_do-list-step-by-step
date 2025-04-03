package usecase

import "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"

func CreateTask(title, description string) domain.Task {
	task := domain.Task{
		ID:          1,
		Title:       title,
		Description: description,
		Done:        false,
	}
	return task
}
