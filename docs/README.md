# betera-test-task

<b>Before starting, insert your Api Key!</b>

### routes

- GET apod/all - all records
- GET apod/{date} - by date (date in format "2023-10-25")
- GET apod/html/{date}

### default ports

- 9001: minio dashboard
- 6555: postgres container
- 8000: http server

### go.mod

- go.uber.org/fx - fx is a dependency injection system for Go.
- github.com/ilyakaznacheev/cleanenv - Minimalistic configuration reader
- log/slog - logger
- gorm.io/gorm - for auto migrations and interaction with postgres
- github.com/minio/minio-go/v7 - minio as s3 storage
- github.com/go-chi/chi/v5 - chi as router

### screenshots

- minio dashboard

<img src="minio-dashboard.png" alt="minio-dashboard"/>

- apod/{date} route

<img src="apod-by-date.png" alt="apod by date route"/>

- apod/all route

<img src="apod-all.png" alt="apod all route"/>

- apod/html/{date} route

<img src="apod-by-date-html.png" alt="apod by date html route"/>