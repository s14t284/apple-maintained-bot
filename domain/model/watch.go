package model

import (
	"time"

	"gorm.io/gorm"
)

// Watch apple watchのドメインオブジェクト
type Watch struct {
	gorm.Model
	Name        string    `gorm:"size:255; not null" json:"name"`
	Strage      string    `gorm:"size:20; not null" json:"strage"`
	Color       string    `gorm:"size:255; not null" json:"color"`
	IsCellular  bool      `gorm:"not null" json:"is_cellular"`
	Amount      int       `gorm:"not null" json:"amount"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
}

// Watches 複数のapple watchのドメインオブジェクト
type Watches []Watch
