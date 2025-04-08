package main

import (
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/interfaces"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/usecase"
)

type fileRepository struct{}

func (f *fileRepository) Save(tasks []domain.Task) error {
	return infrastructure.SaveTasks(tasks)
}

func (f *fileRepository) Load() ([]domain.Task, error) {
	return infrastructure.LoadTasks()
}

func main() {
	// ایجاد نمونه از repository
	repo := &fileRepository{}
	// ایجاد usecase با تزریق repository
	useCase := usecase.NewTaskUseCase(repo)
	// اجرای برنامه CLI
	interfaces.RunCLI(useCase)
}
