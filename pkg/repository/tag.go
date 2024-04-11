package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepository struct {
	collection *mongo.Collection
}

// NewTagRepository create new instance
func NewTagRepository(collection *mongo.Collection) *TagRepository {
	return &TagRepository{
		collection: collection,
	}
}

// NewTagRepository create a new instance
func (t *TagRepository) CreateTag(ctx context.Context, tag *domain.Tag) error {
	_, err := t.collection.InsertOne(ctx, tag)
	if err != nil {
		return err
	}

	return nil
}
