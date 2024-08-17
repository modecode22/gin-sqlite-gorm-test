package routes

import (
	"gin-sqlite-gorm-test/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)

	return r
}