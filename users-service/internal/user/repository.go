package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(user User) error
	DeleteUserByID(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {

	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {

	err := r.db.Create(&user).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ErrDublicateEmail
	}

	return err
}

func (r *userRepository) GetAllUsers() ([]User, error) {

	var users []User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}

func (r *userRepository) GetUserByID(id string) (User, error) {

	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user User) error {

	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUserByID(id string) error {

	return r.db.Delete(&User{}, "id = ?", id).Error
}
