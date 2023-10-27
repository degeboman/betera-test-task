package mapping

import "github.com/degeboman/betera-test-task/internal/models"

// models mapping

func ApodGormToCore(apodGorm models.ApodGorm) models.ApodCore {
	return models.ApodCore{
		Date:           apodGorm.Date,
		Explanation:    apodGorm.Explanation,
		HDUrl:          apodGorm.HDUrl,
		MediaType:      apodGorm.MediaType,
		ServiceVersion: apodGorm.ServiceVersion,
		Title:          apodGorm.Title,
		Url:            apodGorm.Url,
		ImageName:      apodGorm.ImageName,
		HDImageName:    apodGorm.HDImageName,
	}
}

func ApodCoreToGorm(apodCore models.ApodCore) models.ApodGorm {
	return models.ApodGorm{
		Date:           apodCore.Date,
		Explanation:    apodCore.Explanation,
		HDUrl:          apodCore.HDUrl,
		MediaType:      apodCore.MediaType,
		ServiceVersion: apodCore.ServiceVersion,
		Title:          apodCore.Title,
		Url:            apodCore.Url,
	}
}

func ApodWebToCore(apodWeb models.ApodWeb) models.ApodCore {
	return models.ApodCore{
		Date:           apodWeb.Date,
		Explanation:    apodWeb.Explanation,
		HDUrl:          apodWeb.HDUrl,
		MediaType:      apodWeb.MediaType,
		ServiceVersion: apodWeb.ServiceVersion,
		Title:          apodWeb.Title,
		Url:            apodWeb.Url,
	}
}

func ApodCoreToWeb(apodCore models.ApodCore) models.ApodWeb {
	return models.ApodWeb{
		Date:           apodCore.Date,
		Explanation:    apodCore.Explanation,
		HDUrl:          apodCore.HDUrl,
		MediaType:      apodCore.MediaType,
		ServiceVersion: apodCore.ServiceVersion,
		Title:          apodCore.Title,
		Url:            apodCore.Url,
		ImageName:      apodCore.ImageName,
		HDImageName:    apodCore.HDImageName,
	}
}
