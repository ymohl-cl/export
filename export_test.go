package export

import (
	"context"
	"testing"
	"time"
)

func TestCSV(t *testing.T) {
	var e Export
	var err error

	request := ExportRequest{
		Devices:     []int64{0, 1, 2, 3},
		From:        time.Now(),
		To:          time.Now(),
		Filename:    "messages.csv",
		Compression: true,
	}

	e.SetMessageRepository(&MockMessageRepository{})
	if err = e.CSV(context.Background(), request); err != nil {
		t.Error(err)
	}
}
