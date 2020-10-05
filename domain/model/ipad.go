package model

import (
	"time"

	"gorm.io/gorm"
)

// IPad ipadのドメインオブジェクト
type IPad struct {
	gorm.Model
	Name        string    `gorm:"size:255; not null" json:"name"`
	Inch        float32   `gorm:"not null" json:"inch"`
	CPU         string    `gorm:"size:50; not null" json:"cpu"`
	Strage      string    `gorm:"size:20; not null" json:"strage"`
	Camera      string    `gorm:"size:50; not null" json:"camera"`
	Color       string    `gorm:"size:15; not null" json:"color"`
	Amount      int       `gorm:"not null" json:"amount"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
	IsSold      bool      `gorm:"not null" json:"is_sold"`
}

// IPads 複数のipadのドメインオブジェクト
type IPads []IPad
