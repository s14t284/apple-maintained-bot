package model

import (
	"time"

	"gorm.io/gorm"
)

// Mac macbookのドメインオブジェクト
type Mac struct {
	gorm.Model
	Name        string    `gorm:"size:255; not null" json:"name"`
	Inch        float64   `gorm:"not null" json:"inch"`
	CPU         string    `gorm:"size:50; not null" json:"cpu"`
	Memory      int       `gorm:"not null" json:"memory"`
	Strage      string    `gorm:"size:20; not null" json:"strage"`
	TouchBar    bool      `gorm:"not null not null" json:"touchBar"`
	Color       string    `gorm:"size:15; not null" json:"color"`
	Amount      int       `gorm:"not null" json:"amount"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
}

// Macs 複数のmacbookのドメインオブジェクト
type Macs []Mac
