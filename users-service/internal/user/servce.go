package user

import (
	"errors"
)

var ErrDublicateEmail = errors.New("email is already exists")

type Service interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUserByID(id string, user User) (User, error)
	DeleteUserByID(id string) error
}

type userService struct {
	repo UserRepository
}

func NewService(repo UserRepository) Service {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user User) (User, error) {

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetUserByID(id string) (User, error) {

	return s.repo.GetUserByID(id)
}

func (s *userService) GetAllUsers() ([]User, error) {

	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUserByID(id string, user User) (User, error) {

	u, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	u.Email = user.Email

	u.Password = user.Password

	if err := s.repo.UpdateUser(u); err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *userService) DeleteUserByID(id string) error {

	return s.repo.DeleteUserByID(id)
}
