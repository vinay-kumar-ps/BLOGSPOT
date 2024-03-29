package domain

type Media struct {
	ID          int    `json:"id"`
	PostID      int    `json:"post_id"`
	CommentID   int    `json:"comment_id"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
}
