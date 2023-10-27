package apod_worker

import (
	"github.com/degeboman/betera-test-task/constant"
	"github.com/degeboman/betera-test-task/internal/logger/sl"
	"github.com/degeboman/betera-test-task/internal/usecase"
	"log/slog"
	"time"
)

type ApodWorker struct {
	*slog.Logger
	usecase.ApodGetByDateUseCase
	usecase.ApodUploadByDateUseCase
	usecase.ImageDownloadFromUrlUseCase
	usecase.ImageSaveUseCase
	usecase.ApodCreateUseCase
}

func (aw ApodWorker) Ticker(stop chan bool) {
	const op = "apod-worker.apod-worker.Ticker"

	aw.Logger.Info("apod worker is running")

	log := aw.Logger.With(
		slog.String("op", op),
	)

	now := nowInApodFormat()

	// checking that there is today's apod recording
	_, err := aw.ApodGetByDateUseCase.Apply(now)
	if err != nil {
		if err.Error() == constant.ErrRecordNotFound {
			// today record not found
			aw.uploadApodAndSave(nowInApodFormat())
		} else {
			log.Error("failed to get apod meta by date", sl.Err(err))
		}
	}

	ticker := time.NewTicker(24 * time.Hour)

	for {
		select {
		case <-ticker.C:
			aw.uploadApodAndSave(nowInApodFormat())
		case <-stop:
			log.Info("stopping apod ticker")
		}
	}
}

func New(
	logger *slog.Logger,
	apodGetByDateUseCase *usecase.ApodGetByDateUseCaseImpl,
	apodUploadByDateUseCase *usecase.ApodUploadByDateUseCaseImpl,
	imageDownloadFromUrlUseCase *usecase.ImageDownloadFromUrlUseCaseImpl,
	imageSaveUseCase *usecase.ImageSaveUseCaseImpl,
	apodCreateUseCase *usecase.ApodCreateUseCaseImpl,
) ApodWorker {
	return ApodWorker{
		Logger:                      logger,
		ApodGetByDateUseCase:        apodGetByDateUseCase,
		ApodUploadByDateUseCase:     apodUploadByDateUseCase,
		ImageDownloadFromUrlUseCase: imageDownloadFromUrlUseCase,
		ImageSaveUseCase:            imageSaveUseCase,
		ApodCreateUseCase:           apodCreateUseCase,
	}
}

// TODO make delegate
func (aw ApodWorker) uploadApodAndSave(date string) {
	const op = "apod-worker.apod-worker.getApodAndSave"

	log := aw.Logger.With(
		slog.String("op", op),
	)

	// get apod meta by date
	apodCore, err := aw.ApodUploadByDateUseCase.Apply(date)
	if err != nil {
		log.Error("failed to get apod meta", sl.Err(err))
	}

	// cannot save video apod
	if apodCore.MediaType == constant.VideoMediaType {
		return
	}

	// download image by url and save in s3 storage

	imageUnit, err := aw.ImageDownloadFromUrlUseCase.Apply(apodCore.Url)
	if err != nil {
		log.Error("failed to download apod image", sl.Err(err))

		return
	}

	imageName, err := aw.ImageSaveUseCase.Apply(imageUnit)
	if err != nil {
		log.Error("failed to save apod image", sl.Err(err))

		return
	}

	imageHdUnit, err := aw.ImageDownloadFromUrlUseCase.Apply(apodCore.HDUrl)
	if err != nil {
		log.Error("failed to download apod image", sl.Err(err))

		return
	}

	hdImageName, err := aw.ImageSaveUseCase.Apply(imageHdUnit)
	if err != nil {
		log.Error("failed to save apod image", sl.Err(err))

		return
	}

	apodCore.ImageName = imageName
	apodCore.HDImageName = hdImageName

	//save apod model in db
	if err := aw.ApodCreateUseCase.Apply(apodCore); err != nil {
		log.Error("failed to create apod model", sl.Err(err))
	}
}

func nowInApodFormat() string {
	return time.Now().Format("2006-01-02")
}
