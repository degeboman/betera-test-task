package models

import (
	"gorm.io/gorm"
	"time"
)

type ApodResponse struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDUrl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

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

func (ag *ApodGorm) FromResponse(ar *ApodResponse) {
	ag.Date = ar.Date
	ag.Explanation = ar.Explanation
	ag.HDUrl = ar.HDUrl
	ag.MediaType = ar.MediaType
	ag.ServiceVersion = ar.ServiceVersion
	ag.Title = ar.Title
	ag.Url = ar.Url
}
