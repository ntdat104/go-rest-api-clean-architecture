package main

import (
	"fmt"

	"github/go-rest-api-clean-architecture/application/service"
	"github/go-rest-api-clean-architecture/domain/repository"
	"github/go-rest-api-clean-architecture/infrastructure"
	"github/go-rest-api-clean-architecture/interface/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.SetMode("release")

	db := infrastructure.InitSqlite()
	rdb := infrastructure.NewRedisClient()

	// Initialize repositories, services, and handlers
	userRepository := repository.NewUserRepository(db, rdb)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Define routes
	r.POST("/api/v1/users", userHandler.CreateUser)
	r.GET("/api/v1/users", userHandler.GetAllUsers)
	r.GET("/api/v1/users/:id", userHandler.GetUserByID)
	r.PUT("/api/v1/users/:id", userHandler.UpdateUser)
	r.DELETE("/api/v1/users/:id", userHandler.DeleteUser)

	// Run the server
	r.Run(fmt.Sprintf(":%d", 8080))
}
