package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//mongoDB implementation of the comment interface
type CommentRepository struct{
	collection *mongo.Collection
}

//New CommentRepository create a new instance 
func NewCommentRepository(collection *mongo.Collection) *CommentRepository{
	return &CommentRepository{
		collection: collection,
	}
}
//Create create a new comment
func (c *CommentRepository) Create(ctx context.Context,comment  *domain.Comment)error{
	_,err :=c.collection.InsertOne(ctx,comment)
	return err
} 

//Delete deletes a command from database

func(c *CommentRepository) Delete(ctx context.Context ,id int )error{
	_,err :=c.collection.DeleteOne(ctx,bson.M{"id":id})
	return err
}

func  (c *CommentRepository) FindByID(ctx context.Context ,id string) (*domain.Comment,error){

	var comment domain.Comment
	err := c.collection.FindOne(ctx ,bson.M{"id":id}).Decode(&comment)
	if err !=nil{
		return nil,err
	}
	return &comment ,nil
}
func (c *CommentRepository) FindByPostID(ctx context.Context,postID string)([]*domain.Comment,error){

	cursor,err :=c.collection.Find(ctx ,bson.M{"post_id":postID})
	if err !=nil{
		return nil,err
	}
	defer cursor.Close(ctx)

	var comments []*domain.Comment
	for cursor.Next(ctx){
		var comment domain.Comment
		if err :=cursor.Decode(&comment);err !=nil{
			return nil,err
		}
		comments=append(comments, &comment)
	}
	if err := cursor.Err();err !=nil{
		return nil,err
	}
	return comments,nil

	
}
//updates an existing comment in database
func (c *CommentRepository) Update(ctx context.Context ,comment *domain.Comment)error{
	_,err :=c.collection.ReplaceOne(ctx , bson.M{"id":comment.ID},comment)
	return err
}