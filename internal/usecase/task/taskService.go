package Task
import "github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"


type TaskService struct {
	taskRepository  TaskRepository
}

NewTaskService(taskRepository  TaskRepository) *TaskService {
	return &TaskService{taskRepository: TaskRepository}
}

//    - user can create a task and assigne 
// CreateTask یک تسک جدید ایجاد کرده و در ذخیره‌سازی می‌نویسد.
func (u *TaskService) CreateTask(title, description string) (domain.Task, error) {
	tasks, err := u.taskRepository.Load()
	if err != nil {
		return domain.Task{}, err
	}

	// یافتن بیشترین ID برای تولید ID جدید
	newID := 1
	for _, t := range tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}

	task := domain.Task{
		ID:          newID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	tasks = append(tasks, task)

	if err := u.taskRepository.Save(tasks); err != nil {
		return domain.Task{}, err
	}

	return task, nil
}


//TODO://     // - user can change the task status to done 
        
	 
	 
	 
	 
//TODO://	   //- user can delete a task 


//TODO:  // user can list task 

	   
// ListTasks لیست کامل تسک‌ها را برمی‌گرداند.
func (u *TaskUseCase) ListTasks() ([]domain.Task, error) {
	return u.repo.Load()
}
