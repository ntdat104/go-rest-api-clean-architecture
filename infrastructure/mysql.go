package infrastructure

import (
	"github/go-rest-api-clean-architecture/domain/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	const DSN = "root:root@tcp(localhost:3306)/simplize_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}
