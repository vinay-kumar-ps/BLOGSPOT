package repository

import (
	"context"
	"log"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowRepository struct {
	collection *mongo.Collection
}

func NewFollowRepository(collection *mongo.Collection) *FollowRepository {
	return &FollowRepository{
		collection: collection,
	}
}

// follow creates a new folloq relationship in mongoDb
func (f *FollowRepository) Follow(ctx context.Context, follow *domain.Follow) error {
	_, err := f.collection.InsertOne(ctx, follow)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//unfollow removes a folow relationship in monogDb
func (f *FollowRepository ) UnFollow(ctx context.Context,followerID,followingID string)error{
	_,err := f.collection.DeleteOne(ctx ,bson.M{"follower_id":followerID,"following_id":followingID})
	return err
}
//checking if a user following another user
func (f *FollowRepository)IsFollowing(ctx context.Context,followerID,followingID string)(bool,error){
	var result domain.Follow
	err :=f.collection.FindOne(ctx,bson.M{"followe_id":followerID,"following_id":followingID}).Decode(&result)
	if err ==mongo.ErrNoDocuments{
		return false,nil //means there is no relationship
	}else if err !=nil{
		return false,err
	}
	return true,nil

}
