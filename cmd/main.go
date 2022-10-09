package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ianneo97/checkbox/pkg/config/db"
	"github.com/ianneo97/checkbox/pkg/tasks"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/config/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	dbHandler := db.Init(dbUrl)

	tasks.RegisterRoutes(router, dbHandler)

	router.Run(port)
}
