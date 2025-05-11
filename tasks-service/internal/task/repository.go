package task

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID string) ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(task Task) error
	DeleteTaskByID(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TaskRepository {

	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) error {

	return r.db.Create(&task).Error
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {

	var tasks []Task

	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil

}

func (r *taskRepository) GetTasksByUserID(userID string) ([]Task, error) {

	var tasks []Task

	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetTaskByID(id string) (Task, error) {

	var task Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *taskRepository) UpdateTask(task Task) error {

	return r.db.Save(&task).Error
	//return r.db.Model(&task).Updates(task).Error
}

func (r *taskRepository) DeleteTaskByID(id string) error {

	return r.db.Model(&Task{ID: id}).Update("is_done", true).Error
}