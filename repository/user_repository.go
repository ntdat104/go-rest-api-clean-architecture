package repository

import (
	"errors"

	"github/go-rest-api-clean-architecture/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindAll(offset, limit int) ([]model.User, int64, error)
	FindByID(id int64) (*model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	// Get the total count of users
	err := r.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	// Get the paginated result
	err = r.db.Offset(offset).Limit(limit).Order("id DESC").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id int64) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return errors.New("user not found")
	}
	return nil
}
