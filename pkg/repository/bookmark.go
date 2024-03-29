package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoRepository is a Mongo implementation of the bookmark interface
type MongoRepository struct{
	collection *mongo.Collection
}
//NewMongoRepository creates a new instance of MongoRepository
func NewMongoRepository(collection *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		collection: collection,
	}
}

// Create creates a new bookmark in the database.
func (r *MongoRepository) Create(ctx context.Context,bookmark *domain.Bookmark)error{
	_,err :=r.collection.InsertOne(ctx,bookmark)
	return err
}


func (r *MongoRepository) Delete(ctx context.Context,id int)error{
	_,err :=r.collection.DeleteOne(ctx,bson.M{"id":id})
	return err
}

// FindByUserID retrieves all bookmarks for a given user ID from the database.
func (r *MongoRepository) FindByUserID(ctx context.Context, UserID string) ([]*domain.Bookmark, error) {
	// Define a filter to find bookmarks by user ID
	filter := bson.M{"user_id": UserID}

	// Perform a find operation on the collection using the filter
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each bookmark
	var bookmarks []*domain.Bookmark
	for cursor.Next(ctx) {
		var bookmark domain.Bookmark
		if err := cursor.Decode(&bookmark); err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, &bookmark)
	}

	// Check for errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (r *MongoRepository) FindByID(ctx context.Context,postID int)([]*domain.Bookmark,error){
	cursor ,err :=r.collection.Find(ctx,bson.M{"post_id":postID})
	if err !=nil{
		return nil,err
	}
	defer cursor.Close(ctx)

	var bookmarks []*domain.Bookmark
	for cursor.Next(ctx){
		var bookmark domain.Bookmark
		if err :=cursor.Decode(&bookmark);err !=nil {
			return nil,err
		}
		bookmarks =append(bookmarks, &bookmark)
	}
	if err :=cursor.Err();err !=nil{
		return nil,err
	}
	return bookmarks,nil


}