package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ianneo97/checkbox/pkg/config/db"
	"github.com/ianneo97/checkbox/pkg/tasks"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigFile("./pkg/config/envs/.env")
	viper.ReadInConfig()

	gin.SetMode(gin.DebugMode)

	db := db.Init()
	r := setupRouter(db)

	r.Run(":8000")
}

func setupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	tasks.RegisterRoutes(router, db)

	return router
}
