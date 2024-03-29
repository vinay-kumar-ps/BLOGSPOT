package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//UserRepository deines the methods for interacting with user data
type UserRepository interface{

	//CreateUser created a new user
      CreateUser(ctx context.Context,user *domain.User)error

	//   GetUserByID retrives the user by id
    GetUserByID(ctx context.Context, id int)(*domain.User,error)

	//GetUserByEmail retrives a user email
	GetUserByEmail(ctx context.Context,email string)(*domain.User,error)

	// GetUserUsername retrives the user by username
	GetUserUsername(ctx context.Context,username string)(*domain.User,error)

	//UpdateUser updates the user information
	UpdateUser(ctx context.Context ,user *domain.User)error

	//DeleteUser delete a user by ID.
	DeleteUser(ctx context.Context ,id int)error
}
