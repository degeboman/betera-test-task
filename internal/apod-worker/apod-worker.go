package apod_worker

import (
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/constant/apod"
	apodclient "github.com/degeboman/betera-test-task/internal/apod-client"
	"github.com/degeboman/betera-test-task/internal/logger/sl"
	"github.com/degeboman/betera-test-task/internal/models"
	"github.com/degeboman/betera-test-task/internal/storage"
	"log/slog"
	"time"
)

type ApodWorker struct {
	*slog.Logger
	*apodclient.ApodClient
	ApodStorage storage.ApodStorage
}

func (aw ApodWorker) Ticker(stop chan bool) {
	const op = "apod-worker.apod-worker.Ticker"

	aw.Logger.Info("apod worker is running")

	log := aw.Logger.With(
		slog.String("op", op),
	)

	now := nowInApodFormat()

	// checking that there is today's apod recording
	_, err := aw.ApodStorage.ByDate(now)
	if err != nil {
		if err.Error() == constant.ErrRecordNotFound {
			// today record not found
			aw.getApodAndSave(now)
		} else {
			log.Error("failed to get apod meta by date", sl.Err(err))
		}
	}

	ticker := time.NewTicker(24 * time.Hour)

	for {
		select {
		case <-ticker.C:
			aw.getApodAndSave(nowInApodFormat())
		case <-stop:
			log.Info("stopping apod ticker")
		}
	}
}

func nowInApodFormat() string {
	return time.Now().Format("2006-01-02")
}

func New(logger *slog.Logger, ac *apodclient.ApodClient, as storage.ApodStorageImpl) ApodWorker {
	return ApodWorker{
		Logger:      logger,
		ApodClient:  ac,
		ApodStorage: as,
	}
}

func (aw ApodWorker) getApodAndSave(date string) {
	const op = "apod-worker.apod-worker.getApodAndSave"

	log := aw.Logger.With(
		slog.String("op", op),
	)

	apodResponse, err := aw.ApodClient.ByDate(date)
	if err != nil {
		log.Error("failed to get apod meta", sl.Err(err))
	}

	// cannot save video apod
	if apodResponse.MediaType == apod.VideoMediaType {
		return
	}

	var apodGorm models.ApodGorm
	apodGorm.FromResponse(apodResponse)

	// download image from url and save in s3 storage
	imageName, err := aw.ApodClient.DownloadImage(apodGorm.Url)
	if err != nil {
		log.Error("failed to download apod image", sl.Err(err))

		return
	}

	hdImageName, err := aw.ApodClient.DownloadImage(apodGorm.HDUrl)
	if err != nil {
		log.Error("failed to download apod image", sl.Err(err))

		return
	}

	apodGorm.ImageName = imageName
	apodGorm.HDImageName = hdImageName

	//save apod model in db
	if err := aw.ApodStorage.Create(apodGorm); err != nil {
		log.Error("failed to create apod model", sl.Err(err))
	}
}
