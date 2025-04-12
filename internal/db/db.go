package db

import (
	"log"
	"task_manager/internal/taskService"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=1103 dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	if err := db.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("Could not migrate: %s", err)
	}

	return db, nil
}
