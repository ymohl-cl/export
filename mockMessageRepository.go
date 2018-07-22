package export

import (
	"context"
	"strconv"
	"time"
)

// MockMessageRepository : _
type MockMessageRepository struct{}

// List provide a messages list to the device given and the constraint date
func (m MockMessageRepository) List(ctx context.Context, device int64, from, to time.Time) (l []Message, err error) {
	for _, m := range BDD {
		var id int
		if id, err = strconv.Atoi(m.ID); err != nil {
			return nil, err
		}

		if int64(id) == device {
			l = append(l, m)
		}
	}
	return l, nil
}

// BDD provide a set test messages
var BDD = []Message{
	Message{
		Content:   "content0",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content1",
		ID:        "1",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content2",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content3",
		ID:        "2",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content4",
		ID:        "3",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content5",
		ID:        "1",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content6",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content7",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content8",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content9",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content10",
		ID:        "3",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content11",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content12",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content13",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content14",
		ID:        "2",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content15",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content161",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content171",
		ID:        "1",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content01",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content11",
		ID:        "1",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content21",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content31",
		ID:        "2",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content41",
		ID:        "3",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content51",
		ID:        "1",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content61",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content71",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content81",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content91",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content101",
		ID:        "3",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content111",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content121",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content131",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content141",
		ID:        "2",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content151",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content116",
		ID:        "0",
		Timestamp: time.Now(),
	},
	Message{
		Content:   "content171",
		ID:        "1",
		Timestamp: time.Now(),
	},
}
