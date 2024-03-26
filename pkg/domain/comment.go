package domain

import "time"

type Comment struct {
	ID     int `json:"id"`
	PostID int  `json:"post_id"`
	Content  string  `json:"content"`
	AuthorID int `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated _at"`

}
