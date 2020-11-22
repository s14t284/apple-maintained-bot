package model

import (
	"time"
)

// Watch apple watchのドメインオブジェクト
type Watch struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `gorm:"size:255; not null" json:"name"`
	Storage     int       `gorm:"not null;" json:"strage"`
	Color       string    `gorm:"size:255; not null" json:"color"`
	IsCellular  bool      `gorm:"not null" json:"is_cellular"`
	Amount      int       `gorm:"not null" json:"amount"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
	IsSold      bool      `gorm:"not null" json:"is_sold"`
	URL         string    `gorm:"unique;not null" json:"url"`
}

// Watches 複数のapple watchのドメインオブジェクト
type Watches []Watch
