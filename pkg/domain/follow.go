package domain

import "time"

type Follow struct{
	ID int  `json:"id"`
	FollowerID int `json:"follower_id"`
	FollowingID int `json:"following_id"`
	FollowedAt time.Time `json:"followed_at"`

}