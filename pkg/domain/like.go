package domain

// Like represents a user liking a post or comment.
type Like struct {
    UserID   string `json:"user_id"`
    EntityID string `json:"entity_id"`
    EntityType string `json:"entity_type"` // "post" or "comment"
}
