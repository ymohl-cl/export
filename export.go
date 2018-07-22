// Package export provide a lib to export a messaging from client which implement interface MessageRepository
package export

import (
	"compress/gzip"
	"context"
	"errors"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

// Export client module
type Export struct {
	messageClient MessageRepository
}

// MessageRepository define interface to give a messaging client
type MessageRepository interface {
	List(ctx context.Context, device int64, from, to time.Time) ([]Message, error)
}

// SetMessageRepository record the repository to get list message
func (e *Export) SetMessageRepository(repo MessageRepository) {
	e.messageClient = repo
}

// CSV exportation from a request
func (e Export) CSV(ctx context.Context, request ExportRequest) (err error) {
	var wg sync.WaitGroup
	var mErr sync.RWMutex
	var mMessage sync.RWMutex
	var messages []Message

	if e.messageClient == nil {
		return errors.New("message repository not define")
	}

	for _, device := range request.Devices {
		wg.Add(1)
		go func(target int64) {
			defer wg.Done()

			var errAsync error
			var result []Message

			if result, errAsync = e.messageClient.List(ctx, target, request.From, request.To); err != nil {
				mErr.Lock()
				err = errors.New(err.Error() + errAsync.Error())
				mErr.Unlock()
			}
			mMessage.Lock()
			messages = append(messages, result...)
			mMessage.Unlock()
		}(device)
	}
	wg.Wait()

	if err = e.writeFile(request.Filename, messages); err != nil {
		return err
	}
	if request.Compression {
		if err = e.compressFile(request.Filename); err != nil {
			return err
		}
	}
	return nil
}

func (e Export) writeFile(filename string, messages []Message) error {
	var err error
	var file *os.File

	if file, err = os.Create(filename); err != nil {
		return err
	}
	defer file.Close()

	if err = gocsv.MarshalFile(&messages, file); err != nil {
		return err
	}
	return nil
}

func (e Export) compressFile(filename string) error {
	var err error
	var file, compressFile *os.File
	var zw *gzip.Writer
	var content []byte

	if file, err = os.Open(filename); err != nil {
		return err
	}
	defer file.Close()
	if content, err = ioutil.ReadAll(file); err != nil {
		return err
	}
	if compressFile, err = os.Create(filename + ".gz"); err != nil {
		return err
	}
	defer compressFile.Close()

	zw = gzip.NewWriter(compressFile)
	if _, err := zw.Write(content); err != nil {
		return err
	}
	if err = zw.Close(); err != nil {
		return err
	}
	return nil
}
