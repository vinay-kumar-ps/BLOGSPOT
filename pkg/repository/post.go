package repository

import "go.mongodb.org/mongo-driver/mongo"

type PostRepository struct{
	collection *mongo.Collection
}

//NewPostRepository crate a new instance  
