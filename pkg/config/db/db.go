package db

import (
	"log"

	"github.com/ianneo97/checkbox/pkg/tasks"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&tasks.Task{})

	return db
}
