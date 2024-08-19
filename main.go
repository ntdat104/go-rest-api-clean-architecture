package main

import (
	"fmt"
	"log"

	"github/go-rest-api-clean-architecture/handler"
	"github/go-rest-api-clean-architecture/model"
	"github/go-rest-api-clean-architecture/repository"
	"github/go-rest-api-clean-architecture/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	PORT = 8080
	DSN  = "root:root@tcp(localhost:3306)/simplize_dev?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	r := gin.Default()
	// gin.SetMode("release")

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Auto Migrate
	db.AutoMigrate(&model.User{})

	// Initialize repositories, services, and handlers
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Define routes
	r.POST("/api/v1/users", userHandler.CreateUser)
	r.GET("/api/v1/users", userHandler.GetAllUsers)
	r.GET("/api/v1/users/:id", userHandler.GetUserByID)
	r.PUT("/api/v1/users/:id", userHandler.UpdateUser)
	r.DELETE("/api/v1/users/:id", userHandler.DeleteUser)

	// Run the server
	r.Run(fmt.Sprintf(":%d", PORT))
}
