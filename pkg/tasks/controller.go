package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/tasks")

	routes.POST("/", h.AddTask)
	routes.GET("/", h.ListTasks)
	routes.GET("/:id", h.GetTask)
	routes.PATCH("/:id", h.UpdateTask)
	routes.DELETE("/:id", h.DeleteTask)
}
