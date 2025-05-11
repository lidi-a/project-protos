package task

import (
	"github.com/google/uuid"
)

type Service interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID string) ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTaskByID(id string, task Task) (Task, error)
	DeleteTaskByID(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) Service {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task Task) (Task, error) {

	task.ID = uuid.NewString()

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {

	return s.repo.GetAllTasks()
}

func (s *taskService) GetTasksByUserID(userID string) ([]Task, error) {

	return s.repo.GetTasksByUserID(userID)
}

func (s *taskService) GetTaskByID(id string) (Task, error) {

	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTaskByID(id string, task Task) (Task, error) {

	t, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	
	t.Title = task.Title
	t.IsDone = task.IsDone

	if task.UserID != "" {
		t.UserID = task.UserID
	}

	if err := s.repo.UpdateTask(t); err != nil {
		return Task{}, err
	}

	return t, nil

}

func (s *taskService) DeleteTaskByID(id string) error {

	return s.repo.DeleteTaskByID(id)
}
