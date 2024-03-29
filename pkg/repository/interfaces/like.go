package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//LikeRepository defines the mehtods for interacting with like data
type LikeRepository interface{

//like create a new like for a post or comment 
Like(ctx context.Context ,like *domain.Like)error

Unlike(ctx context.Context ,userID ,entityID string)error

//checks if a user has lliked a post or comment
IsLiked(ctx context.Context,userID ,entityID string)error
}