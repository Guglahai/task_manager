package userService

import (
	"task_manager/internal/taskService"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(userReq User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	GetTasksForUser(userID string) ([]taskService.Task, error)
	UpdateUser(id string, userReq User) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(userReq User) (User, error) {
	user := User{
		ID:        uuid.New(),
		Email:     userReq.Email,
		Password:  userReq.Password,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id string) (User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetTasksForUser(userID string) ([]taskService.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *userService) UpdateUser(id string, userReq User) (User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	user.Email = userReq.Email
	user.Password = userReq.Password
	user.UpdatedAt = time.Now()

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
