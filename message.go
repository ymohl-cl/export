package export

import "time"

// Message format
type Message struct {
	ID        string    `csv:"Message.ID"`
	Timestamp time.Time `csv:"Message.Timestamp"`
	Content   string    `csv:"Message.Content"`
}
