package domain

import "time"

type Notification struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	Content  string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ReadStatus bool `json:"read_Status"`
}
