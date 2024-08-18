package services

import (
	"github.com/vshulcz/gopass/internal/crypto"
	"github.com/vshulcz/gopass/internal/models"
	"github.com/vshulcz/gopass/internal/storage"
)

type PasswordService struct {
	storage   storage.Storage
	masterKey []byte
}

func NewPasswordService(storage storage.Storage, masterKey []byte) *PasswordService {
	return &PasswordService{storage: storage, masterKey: masterKey}
}

func (s *PasswordService) AddPassword(service, username, password string) error {
	encryptedPassword, err := crypto.Encrypt(s.masterKey, password)
	if err != nil {
		return err
	}

	entry := models.PasswordEntry{
		Service:  service,
		Username: username,
		Password: encryptedPassword,
	}

	return s.storage.AddEntry(entry)
}

func (s *PasswordService) GetPassword(service string) (string, string, error) {
	entry, err := s.storage.GetEntry(service)
	if err != nil {
		return "", "", err
	}

	decryptedPassword, err := crypto.Decrypt(s.masterKey, entry.Password)
	if err != nil {
		return "", "", err
	}

	return entry.Username, decryptedPassword, nil
}

func (s *PasswordService) DeletePassword(service string) error {
	return s.storage.DeleteEntry(service)
}

func (s *PasswordService) ListPasswords() ([]models.PasswordEntry, error) {
	return s.storage.ListEntries()
}
