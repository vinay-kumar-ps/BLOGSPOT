package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
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
func(m *MediaRepository) GetMediaByID(ctx context.Context ,id interface{} )(*domain.Media ,error){
	var media domain.Media
	filter :=bson.M{"_id":id} //filterttion 
err := m.collection.FindOne(ctx ,filter).Decode(&media)
if err!=nil{
	return nil,err
}
return &media,nil
}
func (m *MediaRepository) DeleteMedia(ctx context.Context, id int)error{
	_,err := m.collection.DeleteOne(ctx,bson.M{"id":id})
	return err
}
