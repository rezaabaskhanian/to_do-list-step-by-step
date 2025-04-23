package Assignee

import (
	"fmt"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/domain"
	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
)

//TODO: struct ro maloom kon va interface ro bede beehsh

// NEW kon azash va address struct jadidi ke mikhad sakhte beshe ro bargardoon

// function haye in struct ro map kon behesh va harja niaz dari repository ro bekhoni azash bekhoon mesle save va load ke dar task estefade kardi

type AssigneeRepositoryFile struct{}

func (r *AssigneeRepositoryFile) Save(tasks []domain.Assignee) error {
	return infrastructure.SaveAssignees(tasks)
}

func (r *AssigneeRepositoryFile) Load() ([]domain.Assignee, error) {
	return infrastructure.LoadAssignees()
}

type AssigneeService struct {
	assigneeRepository AssigneeRepository
}

func NewAssigneeService(assigneeRepository AssigneeRepository) *AssigneeService {
	return &AssigneeService{assigneeRepository: assigneeRepository}
}

// CreateAssignee یک مسئول جدید ایجاد کرده و در ذخیره‌سازی می‌نویسد.
func (u *AssigneeService) CreateAssignee(name, email string) (domain.Assignee, error) {
	assignees, err := u.assigneeRepository.Load()
	if err != nil {
		return domain.Assignee{}, err
	}

	// یافتن بیشترین ID برای تولید ID جدید
	newID := 1
	for _, a := range assignees {
		if a.ID >= newID {
			newID = a.ID + 1
		}
	}

	assignee := domain.Assignee{
		ID:    newID,
		Name:  name,
		Email: email,
	}

	assignees = append(assignees, assignee)

	if err := u.assigneeRepository.Save(assignees); err != nil {
		return domain.Assignee{}, err
	}

	return assignee, nil
}

// ListAssignees لیست کامل مسئولین را برمی‌گرداند.
func (u *AssigneeService) ListAssignees() ([]domain.Assignee, error) {
	return u.assigneeRepository.Load()
}

// UpdateAssignee اطلاعات یک مسئول را به‌روز می‌کند.
func (u *AssigneeService) UpdateAssignee(id int, name, email string) (domain.Assignee, error) {
	assignees, err := u.assigneeRepository.Load()
	if err != nil {
		return domain.Assignee{}, err
	}

	var updatedAssignee domain.Assignee
	found := false
	for i := range assignees {
		if assignees[i].ID == id {
			assignees[i].Name = name
			assignees[i].Email = email
			updatedAssignee = assignees[i]
			found = true
			break
		}
	}

	if !found {
		return domain.Assignee{}, fmt.Errorf("مسئول با ID %d یافت نشد", id)
	}

	if err := u.assigneeRepository.Save(assignees); err != nil {
		return domain.Assignee{}, err
	}

	return updatedAssignee, nil
}

// DeleteAssignee یک مسئول را از لیست حذف می‌کند.
func (u *AssigneeService) DeleteAssignee(id int) error {
	assignees, err := u.assigneeRepository.Load()
	if err != nil {
		return err
	}

	newAssignees := make([]domain.Assignee, 0, len(assignees))
	found := false
	for _, a := range assignees {
		if a.ID == id {
			found = true
			continue
		}
		newAssignees = append(newAssignees, a)
	}

	if !found {
		return fmt.Errorf("مسئول با ID %d یافت نشد", id)
	}

	return u.assigneeRepository.Save(newAssignees)
}
