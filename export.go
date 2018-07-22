// Package export provide a lib to export a messaging from client which implement interface MessageRepository
package export

import (
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
)

// Export client module
type Export struct {
	messageClient MessageRepository
	csvFile       *os.File
	compressFile  *os.File
	zw            *gzip.Writer
	mFile         sync.RWMutex
	mCompress     sync.RWMutex
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
func (e *Export) CSV(ctx context.Context, request ExportRequest) (err error) {
	var wg sync.WaitGroup
	var m sync.RWMutex

	if e.messageClient == nil {
		return errors.New("message repository not define")
	}
	if err = e.initCSV(request.Filename); err != nil {
		return err
	}
	defer e.closeCSV()
	if request.Compression {
		if err = e.initCompress(request.Filename); err != nil {
			return err
		}
		defer e.closeCompress()
	}

	for _, device := range request.Devices {
		wg.Add(1)
		go func(target int64) {
			defer wg.Done()
			var errAsync error
			var result []Message

			m.Lock()
			if result, errAsync = e.messageClient.List(ctx, target, request.From, request.To); err != nil {
				fmt.Printf("error: %s, to access message from bdd client to the device: %d\n", errAsync.Error(), target)
				m.Unlock()
				return
			}
			m.Unlock()
			if errAsync = e.writeMessage(result); err != nil {
				fmt.Printf("error: %s, can't write message list to the device: %d\n", errAsync.Error(), target)
				return
			}
			if request.Compression {
				if errAsync = e.compressMessage(result); err != nil {
					fmt.Printf("error compression messages: %s, to the device: %d\n", errAsync.Error(), target)
					return
				}
			}
		}(device)
	}
	wg.Wait()
	return nil
}

func (e *Export) initCSV(filename string) error {
	var err error

	if e.csvFile, err = os.Create(filename); err != nil {
		return err
	}
	return nil
}

func (e *Export) initCompress(filename string) error {
	var err error

	if e.compressFile, err = os.Create(filename + ".gz"); err != nil {
		return err
	}
	e.zw = gzip.NewWriter(e.compressFile)
	return nil
}

func (e *Export) closeCSV() error {
	var err error

	if err = e.csvFile.Close(); err != nil {
		return err
	}
	return nil
}

func (e *Export) closeCompress() error {
	var err error

	if err = e.zw.Close(); err != nil {
		return err
	}
	if err = e.compressFile.Close(); err != nil {
		return err
	}
	return nil
}

func (e *Export) writeMessage(messages []Message) error {
	var err error

	e.mFile.Lock()
	defer e.mFile.Unlock()
	if err = gocsv.MarshalFile(&messages, e.csvFile); err != nil {
		return err
	}
	return nil
}

func (e *Export) compressMessage(messages []Message) error {
	var err error
	var buf []byte

	if buf, err = gocsv.MarshalBytes(&messages); err != nil {
		return err
	}
	e.mCompress.Lock()
	defer e.mCompress.Unlock()
	if _, err := e.zw.Write(buf); err != nil {
		return err
	}
	return nil
}
