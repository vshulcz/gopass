package storage

import (
	"github.com/vshulcz/gopass/internal/models"

	"gorm.io/gorm"
)

type Storage interface {
	AddEntry(entry models.PasswordEntry) error
	GetEntry(service string) (*models.PasswordEntry, error)
	DeleteEntry(service string) error
	ListEntries() ([]models.PasswordEntry, error)
}

type GormStorage struct {
	db *gorm.DB
}

func NewGormStorage(db *gorm.DB) *GormStorage {
	return &GormStorage{db: db}
}

func (s *GormStorage) AddEntry(entry models.PasswordEntry) error {
	return s.db.Create(&entry).Error
}

func (s *GormStorage) GetEntry(service string) (*models.PasswordEntry, error) {
	var entry models.PasswordEntry
	err := s.db.Where("service = ?", service).First(&entry).Error
	return &entry, err
}

func (s *GormStorage) DeleteEntry(service string) error {
	return s.db.Where("service = ?", service).Delete(&models.PasswordEntry{}).Error
}

func (s *GormStorage) ListEntries() ([]models.PasswordEntry, error) {
	var entries []models.PasswordEntry
	err := s.db.Find(&entries).Error
	return entries, err
}
