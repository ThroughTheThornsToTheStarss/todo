package postgress

import (
	"gorm.io/gorm"
	repo "github.com/ThroughTheThornsToTheStarss/todo/internal/repo/postgress"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&repo.Todo{})
}