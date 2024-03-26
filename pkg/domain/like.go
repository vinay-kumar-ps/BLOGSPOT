package domain

import "time"

type Like struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	CommentID int       `json:"comment_id"`
	LikedAT   time.Time `json:"liked_at"`
}
