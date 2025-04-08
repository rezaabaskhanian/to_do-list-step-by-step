package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
)

const storageFile = "tasks.json"

// SaveTasks کل لیست تسک‌ها را به صورت JSON در فایل ذخیره می‌کند.
func SaveTasks(tasks []domain.Task) error {
	file, err := os.OpenFile(storageFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("خطا در باز کردن یا ایجاد فایل: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("خطا در کدگذاری JSON: %w", err)
	}

	return nil
}

// LoadTasks تسک‌ها را به صورت آرایه از فایل بارگذاری می‌کند.
func LoadTasks() ([]domain.Task, error) {
	file, err := os.Open(storageFile)
	if err != nil {
		// اگر فایل وجود نداشت، آرایه‌ی خالی برگردان
		if os.IsNotExist(err) {
			return []domain.Task{}, nil
		}
		return nil, fmt.Errorf("خطا در باز کردن فایل: %w", err)
	}
	defer file.Close()

	var tasks []domain.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("خطا در دیکد کردن JSON: %w", err)
	}
	return tasks, nil
}
