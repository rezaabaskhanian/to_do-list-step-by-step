package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
)

// ذخیره‌سازی تسک‌ها در فایل
const storageFile = "tasks.json"

func SaveTasks(tasks []domain.Task) error {
	// باز کردن فایل برای افزودن داده‌ها (ایجاد در صورت عدم وجود)
	file, err := os.OpenFile(storageFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("خطا در باز کردن یا ایجاد فایل: %w", err)
	}
	defer file.Close()

	// ایجاد انکودر JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	// نوشتن داده‌های جدید به فایل
	for _, task := range tasks {
		if err := encoder.Encode(task); err != nil {
			return fmt.Errorf("خطا در کدگذاری JSON: %w", err)
		}
	}
	return nil
}

// بازیابی تسک‌ها از فایل
func LoadTasks() ([]domain.Task, error) {
	file, err := os.Open(storageFile)
	if err != nil {
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
