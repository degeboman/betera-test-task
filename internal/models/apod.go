package models

import (
	"gorm.io/gorm"
	"time"
)

// ApodWeb structure for interaction via http
type ApodWeb struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDUrl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
	ImageName      string `json:"image_name"`
	HDImageName    string `json:"hd_image_name"`
}

// ApodCore is implementation independent
type ApodCore struct {
	Date           string
	Explanation    string
	HDUrl          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
	ImageName      string
	HDImageName    string
}

// ApodGorm is special model for gorm
type ApodGorm struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Date           string         `gorm:"not null;"`
	Explanation    string
	HDUrl          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
	ImageName      string `gorm:"not null"`
	HDImageName    string `gorm:"not null"`
}
