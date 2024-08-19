package service

import (
	"time"

	"github/go-rest-api-clean-architecture/model"
	"github/go-rest-api-clean-architecture/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetAllUsers(offset, limit int) ([]model.User, int64, error)
	GetUserByID(id int64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(user *model.User) error {
	now := time.Now()
	user.CreatedDate = &now
	user.UpdatedDate = &now
	return s.userRepository.Create(user)
}

func (s *userService) GetAllUsers(offset, limit int) ([]model.User, int64, error) {
	return s.userRepository.FindAll(offset, limit)
}

func (s *userService) GetUserByID(id int64) (*model.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	user.UpdatedDate = new(time.Time)
	*user.UpdatedDate = time.Now()
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.userRepository.Delete(id)
}
