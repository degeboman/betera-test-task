package apod_worker

import (
	"testing"
	"time"
)

func TestApodWorker_nowInApodFormat(t *testing.T) {
	result := nowInApodFormat()

	if result != "2023-10-24" {
		t.Error("incorrect result: expected 2023-10-24, got: ", result)
	}
}

func TestApodWorker_Ticker(t *testing.T) {
	aw := ApodWorker{}

	stop := make(chan bool)

	go aw.Ticker(stop)

	time.Sleep(10 * time.Hour)
	stop <- true
}
