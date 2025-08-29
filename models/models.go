package models

import "time"

// User model
type User struct {
	ID           int       `gorm:"primaryKey"`
	Username     string    `gorm:"not null;unique"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

// Terminal model
type Terminal struct {
	ID        int       `gorm:"primaryKey"`
	Code      string    `gorm:"not null;unique"`
	Name      string    `gorm:"not null"`
	Location  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
