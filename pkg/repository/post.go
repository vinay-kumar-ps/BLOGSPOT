package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct{
	collection *mongo.Collection
}

//NewPostRepository crate a new instance  
func NewPostRepository(collection *mongo.Collection) *PostRepository{
	return &PostRepository{
		collection: collection,
	}
}
//this will create a new post
func(p *PostRepository) Create(post *domain.Post)error{
	_,err :=p .collection.InsertOne(context.Background(),post)
	if err !=nil{
		return err 
	}
	return nil
}
// FindByID method finds a post by ID
func (p *PostRepository) FindByUserID(userID string) ([]*domain.Post, error) {
	cursor, err := p.collection.Find(context.Background(), bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []*domain.Post
	for cursor.Next(context.Background()) {
		var post domain.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
func(p *PostRepository)Update(id string,post *domain.Post)error{
	_,err :=p.collection.ReplaceOne(context.Background(),bson.M{"id":id},post)
	if err !=nil{
		return err
	}
	return nil
}
