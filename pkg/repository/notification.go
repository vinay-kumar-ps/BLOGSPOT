package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository struct{
	collection *mongo.Collection
}

// NewNotificationRepository create a new instance
func NewNotificationRepository( collection *mongo.Collection) *NotificationRepository{
	return &NotificationRepository{
		collection: collection,
	}
	
}
//this will create a new notification 
func(n *NotificationRepository) CreateNotification(ctx context.Context, notification *domain.Notification)error{
	_,err := n.collection.InsertOne(ctx,notification)
	return err
}

// GetNotificationsByUserID  get notifications to the  users
func (n *NotificationRepository) GetNotificationsByUserID(ctx context.Context,userID int)([]*domain.Notification,error){
	cursor,err := n.collection.Find(ctx, bson.M{"user_id":userID})
	if err !=nil{
		return nil,err
	}
	defer cursor.Close(ctx)

	var notifications []*domain.Notification
	for cursor.Next(ctx){
		var notification domain.Notification
		if err :=cursor.Decode(&notification);err !=nil{
			return nil,err
		}
		notifications =append(notifications,&notification)
	}
	if err :=cursor.Err();err!=nil{
		return nil,err
	}
	return notifications,nil

}

func(n *NotificationRepository)MarkNotificationAsRead(ctx context.Context,id int)error{
	_,err :=n.collection.UpdateOne(
		ctx,
		bson.M{"id":id},
		bson.M{"set": bson.M{"read":true}},
	)
	return err
}