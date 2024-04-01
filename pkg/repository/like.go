package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository struct{
	collection *mongo.Collection
}

//crating a new instance of likerepository
func NewLikeRepository( collection *mongo.Collection) *LikeRepository{
	return &LikeRepository{
		collection: collection,
	}
}
//creates like for a post
func (l *LikeRepository) Like(ctx context.Context,like *domain.Like)error{
	_,err := l.collection.InsertOne(ctx,like)
	return err
	
}
// removes a like for the post 
func (l *LikeRepository) Unlike(ctx context.Context,userID,entityID string )error{
	_,err :=l.collection.DeleteOne(ctx,bson.M{"user_id":userID,"entity_id":entityID})
	return err
}

//checking if a user has liked a post 
func( l *LikeRepository) IsLiked(ctx context.Context,userID,entityId string)(bool,error){
	count,err :=l.collection.CountDocuments(ctx,bson.M{"user_id":userID,"entity_id":entityId})
	if err !=nil{
		return false,err
	}
	return count > 0 ,nil
}