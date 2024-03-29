package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//TagRepository  defines the methods for interacting with tag data
type TagRepository interface{
	//creating a new tag
	CreateTag(ctx context.Context ,tag *domain.Tag )error

	// GetTagByID retrives tags by ID
      GetTagByID(ctx context.Context ,id int)(*domain.Tag,error)

	  //GetTagByName retrives a tag by name 
	  GetTagByName(ctx context.Context , name string)(domain.Tag,error)

	//   GetAllTags retrives all tags 
	GetAllTags(ctx context.Context )([]*domain.Tag,error)

	// DeleteTag deletes a tag by ID
    DeleteTag(ctx context.Context,id int)error

	} 