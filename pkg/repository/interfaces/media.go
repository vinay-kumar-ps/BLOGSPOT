package repository

import (
	"context"

	"github.com/vinay-kumar-ps/blogspot/pkg/domain"
)

//defines the mehtods for interacting  with media data.
type MediaRepository interface{

	//uploads new media fiiles 
	UploadMedia(ctx context.Context,media *domain.Media)error

	//getmediabyid retrives a media file by id 
	GetMediaByID(ctx context.Context,id int ) (*domain.Media,error)

	//DeleteMedia deltes a media file by id .
	DeleteMedia(ctx context.Context,id int )error
}