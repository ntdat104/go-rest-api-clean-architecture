package repository

import (
	"errors"
	"time"

	"github/go-rest-api-clean-architecture/domain/model"
	"github/go-rest-api-clean-architecture/infrastructure"
	"github/go-rest-api-clean-architecture/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindAll(offset, limit int) (*[]model.User, int64, error)
	FindByID(id int64) (*model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userRepository struct {
	db     *gorm.DB
	rdb    *infrastructure.RedisClient
	prefix string
}

func NewUserRepository(db *gorm.DB, rdb *infrastructure.RedisClient) UserRepository {
	return &userRepository{db: db, rdb: rdb, prefix: "users"}
}

func (r *userRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	r.rdb.InvalidPrefix(r.prefix)
	return nil
}

func (r *userRepository) FindAll(offset, limit int) (*[]model.User, int64, error) {
	key := utils.GenerateKey(r.prefix, offset, limit)
	totalKey := utils.GenerateKey(r.prefix, "total")
	const TTL = 3 * time.Minute

	var users []model.User
	var total int64

	if exists, _ := r.rdb.Exists(key); exists {
		r.rdb.GetMsgPack(key, &users)
		r.rdb.GetMsgPack(totalKey, &total)
		return &users, total, nil
	}

	err := r.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset(offset).Limit(limit).Order("id DESC").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	r.rdb.SetMsgPack(key, users, TTL)
	r.rdb.SetMsgPack(totalKey, total, TTL)

	return &users, total, nil
}

func (r *userRepository) FindByID(id int64) (*model.User, error) {
	key := utils.GenerateKey(r.prefix, id)
	const TTL = 3 * time.Minute

	var user *model.User
	r.rdb.GetMsgPack(key, user)
	if user != nil {
		return user, nil
	}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	r.rdb.SetMsgPack(key, user, TTL)
	return user, nil
}

func (r *userRepository) Update(user *model.User) error {
	key := utils.GenerateKey(r.prefix, user.Id)
	if err := r.db.Save(user).Error; err != nil {
		return err
	}
	r.rdb.Invalidate(key)
	return nil
}

func (r *userRepository) Delete(id int64) error {
	key := utils.GenerateKey(r.prefix, id)
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return errors.New("user not found")
	}
	r.rdb.Invalidate(key)
	return nil
}
