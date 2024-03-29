package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//postRepository defines the methods for interacting  with post data.
type PostRepository interface {
	FindByID(ctx context.Context,id string) (*domain.Post,error)
	FIndByUserID(ctx context.Context,userID string) ([]*domain.Post,error)
	Create(ctx context.Context,post *domain.Post)error
	Update(ctx context.Context,id string)error

}