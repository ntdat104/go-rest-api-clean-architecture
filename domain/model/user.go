package model

import "time"

type User struct {
	Id          int64      `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	FullName    *string    `json:"fullName" gorm:"column:full_name"`
	Email       *string    `json:"email" gorm:"column:email"`
	PhoneNumber *string    `json:"phoneNumber" gorm:"column:phone_number"`
	IsMale      *bool      `json:"isMale" gorm:"column:is_male"`
	Status      *int       `json:"status" gorm:"column:status"`
	CreatedDate *time.Time `json:"createdDate" gorm:"column:created_date"`
	UpdatedDate *time.Time `json:"updatedDate" gorm:"column:updated_date"`
}
