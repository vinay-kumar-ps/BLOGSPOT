package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//CommentRepository  defines the methods for interacting with comment data
type  CommentRepository interface{
	
	//create creates a new comment .
	Create(ctx context.Context,comment*domain.Comment)error


	//Delete deletes a comment  by ID.
	 Delete( ctx context.Context,id string)error
	 

	//FindByID retrives a comment by ID.
	FindByID(ctx context.Context,id string) (*domain.Comment,error)

	//FindByUserPostID retrives all comments for a given post ID.
	FindByPostID(ctx context.Context,postID string) ([]*domain.Comment,error)

	

	//update updates an existing comment .
	Update(ctx context.Context,comment *domain.Comment)error


	
}