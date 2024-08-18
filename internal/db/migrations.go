package db

import (
	"github.com/vshulcz/gopass/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.PasswordEntry{})
}
