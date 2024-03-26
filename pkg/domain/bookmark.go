package domain

import "time"

type Bookmark struct{
	ID  int `json:"id"`
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
	BookmarkedAt  time.Time `json:"bookmarked_at"`
}