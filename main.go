package main

import (
	"fmt"
	"log"

	"github/go-rest-api-clean-architecture/handler"
	"github/go-rest-api-clean-architecture/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	PORT = 8080
	DSN  = "root:root@tcp(127.0.0.1:3306)/simplize_dev?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Auto Migrate
	db.AutoMigrate(&model.User{})

	// Initialize repositories, services, and handlers
	userHandler := handler.NewUserHandler(db)

	// Define routes
	r.POST("/api/v1/users", userHandler.CreateUser)
	r.GET("/api/v1/users", userHandler.GetAllUsers)
	r.GET("/api/v1/users/:id", userHandler.GetUserByID)
	r.PUT("/api/v1/users/:id", userHandler.UpdateUser)
	r.DELETE("/api/v1/users/:id", userHandler.DeleteUser)

	// Run the server
	r.Run(fmt.Sprintf(":%d", PORT))
}
