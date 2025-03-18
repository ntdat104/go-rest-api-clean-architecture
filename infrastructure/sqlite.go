// InitSqlite
package infrastructure

import (
	"fmt"
	"github/go-rest-api-clean-architecture/domain/model"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.AutoMigrate(&model.User{})
	now := time.Now()
	for i := range make([]struct{}, 100000) {
		fullname := fmt.Sprintf("user_%d", i+1)
		email := fmt.Sprintf("user_%d@gmail.com", i+1)
		phoneNumber := fmt.Sprintf("098765%04d", i+1) // Example phone format
		isMale := true
		status := 1

		user := model.User{
			FullName:    &fullname,
			Email:       &email,
			PhoneNumber: &phoneNumber,
			IsMale:      &isMale,
			Status:      &status,
			CreatedDate: &now,
			UpdatedDate: &now,
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Println("Error inserting user:", err)
		}
	}
	return db
}
