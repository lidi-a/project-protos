package database

import (
	"log"
	"os"

	"github.com/lidi-a/project-protos/tasks-service/internal/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRESDSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	return db, nil
}
