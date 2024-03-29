package domain

// Follow represents a user following another user.
type Follow struct {
    FollowerID string `json:"follower_id"`
    FollowingID string `json:"following_id"`
}