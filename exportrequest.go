package export

import "time"

// ExportRequest is a model to request messages list
// call it Request to golint
type ExportRequest struct {
	Devices     []int64
	From, To    time.Time
	Filename    string
	Compression bool
}
