package userService

import (
	"task_manager/internal/taskService"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) error
	GetAllUsers() ([]User, error)
	GetTasksByUserID(id string) ([]taskService.Task, error)
	GetUserByID(id string) (User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetTasksByUserID(id string) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := r.db.Find(&tasks, "user_id = ?", id).Error
	return tasks, err
}

func (r *userRepository) GetUserByID(id string) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
