package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MediaRepository struct {
	collection *mongo.Collection
}
func NewMediaRepository(collection *mongo.Collection) *MediaRepository{
	return &MediaRepository{
		collection: collection,
	}
}
//  UploadMedia create a new media 
func (m *MediaRepository) UploadMedia(ctx context.Context ,media interface{} )error{
	_,err := m.collection.InsertOne(ctx ,media)
	return err
}
// GetMediaByID get the  specified media 
func(m *MediaRepository) GetMediaByID(ctx context.Context ,media )
_,err :=m.collection