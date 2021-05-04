package model

import (
	"time"
)

// Mac macbookのドメインオブジェクト
type Mac struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `gorm:"size:255; not null" json:"name"`
	Inch        float32   `gorm:"not null" json:"inch"`
	CPU         string    `gorm:"size:50; not null" json:"cpu"`
	Memory      int       `gorm:"not null" json:"memory"`
	Storage     int       `gorm:"not null;" json:"strage"`
	TouchBar    bool      `gorm:"not null not null" json:"touch_bar"`
	Color       string    `gorm:"size:15; not null" json:"color"`
	Amount      int       `gorm:"not null" json:"amount"`
	ReleaseDate time.Time `gorm:"not null" json:"release_date"`
	IsSold      bool      `gorm:"not null" json:"is_sold"`
	URL         string    `gorm:"unique;not null" json:"url"`
}

// Macs 複数のmacbookのドメインオブジェクト
type Macs []Mac

// MacRequestParam macを検索するときのリクエストパラメータを格納するオブジェクト
type MacRequestParam struct {
	Name       string
	Color      string
	TouchBar   string
	IsSold     string
	MaxInch    float64
	MinInch    float64
	MaxMemory  int
	MinMemory  int
	MaxStorage int
	MinStorage int
	MaxAmount  int
	MinAmount  int
}
