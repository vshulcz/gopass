package models

type PasswordEntry struct {
	ID       uint `gorm:"primaryKey"`
	Service  string
	Username string
	Password string
}
