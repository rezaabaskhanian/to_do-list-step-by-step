package usecase

import "/internal/domain"

func CreateTask(title, description string) domain.Task {
	task := domain.Task{
		ID:          1,
		Title:       title,
		Description: description,
		Done:        false,
	}
	return task
}
