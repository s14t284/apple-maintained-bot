package model

import (
	"time"
)

// IPad ipadのドメインオブジェクト
type IPad struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `gorm:"size:255; not null" json:"name"`
	Inch        float32   `gorm:"not null" json:"inch"`
	Storage     int       `gorm:"not null;" json:"strage"`
	Camera      string    `gorm:"size:50; not null" json:"camera"`
	Color       string    `gorm:"size:15; not null" json:"color"`
	Amount      int       `gorm:"not null" json:"amount"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
	IsSold      bool      `gorm:"not null" json:"is_sold"`
	URL         string    `gorm:"unique;not null" json:"url"`
}

// IPads 複数のipadのドメインオブジェクト
type IPads []IPad

// IPadRequestParam ipadを検索するときのリクエストパラメータを格納するオブジェクト
type IPadRequestParam struct {
	Name       string
	Color      string
	IsSold     string
	MaxInch    float64
	MinInch    float64
	MaxStorage int
	MinStorage int
	MaxAmount  int
	MinAmount  int
}
