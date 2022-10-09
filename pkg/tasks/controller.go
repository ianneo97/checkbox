package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}

	routes := router.Group("/tasks")

	routes.POST("/", h.AddTask)
	routes.GET("/all", h.ListTasks)
	routes.GET("/:id", h.GetTask)
	routes.PATCH("/:id", h.UpdateTask)
	routes.DELETE("/:id", h.DeleteTask)
}
