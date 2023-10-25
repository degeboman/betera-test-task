package models

import (
	"gorm.io/gorm"
	"time"
)

type ApodCore struct {
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HDUrl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
	ImageName      string
	HDImageName    string
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

func (ag *ApodGorm) ToCore() ApodCore {
	return ApodCore{
		Date:           ag.Date,
		Explanation:    ag.Explanation,
		HDUrl:          ag.HDUrl,
		MediaType:      ag.MediaType,
		ServiceVersion: ag.ServiceVersion,
		Title:          ag.Title,
		Url:            ag.Url,
		ImageName:      ag.ImageName,
		HDImageName:    ag.HDImageName,
	}
}

func (ag *ApodGorm) FromCore(ac *ApodCore) {
	ag.Date = ac.Date
	ag.Explanation = ac.Explanation
	ag.HDUrl = ac.HDUrl
	ag.MediaType = ac.MediaType
	ag.ServiceVersion = ac.ServiceVersion
	ag.Title = ac.Title
	ag.Url = ac.Url
}
