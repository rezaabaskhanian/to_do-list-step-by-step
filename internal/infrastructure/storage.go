package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
)

const taskStorageFile = "tasks.json"
const assigneeStorageFile = "assignees.json"

// InitializeStorage بررسی می‌کنه که فایل‌های ذخیره‌سازی وجود دارند یا نه. اگر نه، ایجادشون می‌کنه.
func InitializeStorage() error {
	files := []string{taskStorageFile, assigneeStorageFile}

	for _, file := range files {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			// اگر فایل وجود نداشت، بسازش
			err := os.MkdirAll("data", os.ModePerm)
			if err != nil {
				return err
			}
			f, err := os.Create(file)
			if err != nil {
				return err
			}
			f.Close()
		} else if err != nil {
			// خطای دیگه‌ای به‌جز وجود نداشتن فایل
			return err
		}
	}

	return nil
}

//TODO bere task.json va assignee.json kham ro besaze hamoo aval . va check kone age vojud dasht dige nasaze

// SaveTasks کل لیست تسک‌ها را به صورت JSON در فایل ذخیره می‌کند.
func SaveTasks(tasks []domain.Task) error {
	file, err := os.OpenFile(taskStorageFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
	file, err := os.Open(taskStorageFile)
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

//TODO : Save Assignee

// SaveAssignees کل لیست Assignee‌ها را در فایل JSON ذخیره می‌کند.
func SaveAssignees(assignees []domain.Assignee) error {
	file, err := os.OpenFile(assigneeStorageFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("خطا در باز کردن یا ایجاد فایل: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(assignees); err != nil {
		return fmt.Errorf("خطا در کدگذاری JSON: %w", err)
	}

	return nil
}

// TODO : Load Assignee

// LoadAssignees لیست Assignee‌ها را از فایل JSON می‌خواند.
func LoadAssignees() ([]domain.Assignee, error) {
	file, err := os.Open(assigneeStorageFile)
	if err != nil {
		// اگر فایل وجود نداشت، آرایه‌ی خالی برگردان
		if os.IsNotExist(err) {
			return []domain.Assignee{}, nil
		}
		return nil, fmt.Errorf("خطا در باز کردن فایل: %w", err)
	}
	defer file.Close()

	var assignees []domain.Assignee
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&assignees); err != nil {
		return nil, fmt.Errorf("خطا در دیکد کردن JSON: %w", err)
	}
	return assignees, nil
}
