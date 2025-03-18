package repository

import (
	"errors"
	"fmt"

	"github/go-rest-api-clean-architecture/domain/model"
	"github/go-rest-api-clean-architecture/infrastructure"

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
	db  *gorm.DB
	rdb *infrastructure.RedisClient
}

func NewUserRepository(db *gorm.DB, rdb *infrastructure.RedisClient) UserRepository {
	return &userRepository{db: db, rdb: rdb}
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	// r.rdb.InvalidPrefix("users:")
	return nil
}

func (r *userRepository) FindAll(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	// err := r.rdb.GetJSON(fmt.Sprintf("users:%d:%d", offset, limit), &users)
	// if err == nil {
	// 	return users, int64(len(users)), nil
	// }

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

	// r.rdb.SetJSON(fmt.Sprintf("users:%d:%d", offset, limit), users, 3*time.Minute)

	return users, total, nil
}

func (r *userRepository) FindByID(id int64) (*model.User, error) {
	var user model.User
	// err := r.rdb.GetJSON(fmt.Sprintf("users:%d", id), &user)
	// if err == nil {
	// 	return &user, nil
	// }
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	// r.rdb.SetJSON(fmt.Sprintf("users:%d", id), user, 3*time.Minute)
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	r.rdb.Invalidate(fmt.Sprintf("users:%d", user.Id))
	return nil
}

func (r *userRepository) Delete(id int64) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return errors.New("user not found")
	}
	r.rdb.Invalidate(fmt.Sprintf("users:%d", id))
	return nil
}
