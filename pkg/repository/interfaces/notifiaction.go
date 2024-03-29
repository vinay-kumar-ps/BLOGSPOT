package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//NotificationRepository defines the mehtods for getting notifications
type NotificationRepository interface{

	//creates a new notification 
	CreateNotification(ctx context.Context,notification *domain.Notification)error

	//GetNotificationsByUserID retrives all notification for a given user ID.
	GetNotificationsByUserID(ctx context.Context,userID int)([]*domain.Notification,error)

	//MarkNotificationAsRead marks a notification as read
	MarkNotificationAsRead(ctx context.Context,id int)error
}