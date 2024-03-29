package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//defines the methods for interacting eith follow data.
type FollowRepository interface{

	//follow creates new follow relationship
	Follow(cts context.Context,follow *domain.Follow)error

	//unfollow reomves a follow relationship
	 UnFollow(ctx context.Context,followerID,followingID string )error

	 //checking if a user is followign another user
	 IsFollowing(ctx context.Context,followerID,followingID string)(bool ,error)

}