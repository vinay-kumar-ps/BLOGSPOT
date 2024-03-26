package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//Bookmarkrepository defines the methods for interacting with bookmark data.
type BookmarkRepository interface{

	//FindByID retrives the bookmark by id.
	FindByID (ctx context.Context ,id string) (*domain.Bookmark,error)

	//FindByUserId retrives all bookmarks for a given userID.
	FIndByUserID(ctx context.Context,UserID string) ([]domain.Bookmark,error) 

	//create  creates a new bookmark 
	Create (ctx context.Context, bookmark *domain.Bookmark) error
	
	// Delete deletes a bookmark by ID
	Delete (ctx context.Context ,id string)error

}