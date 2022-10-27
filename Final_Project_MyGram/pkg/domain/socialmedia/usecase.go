package socialmedia

import (
	"context"

	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/message"
)

type SocialMediaUsecase interface {
	CreateSocialMediaSvc(ctx context.Context, socialMedia SocialMedia) (result SocialMedia, errMsg message.ErrorMessage)
	GetSocialMediasSvc(ctx context.Context) (result []SocialMedia, errMsg message.ErrorMessage)
	GetSocialMediaByIdSvc(ctx context.Context, socmedId uint64) (result SocialMedia, errMsg message.ErrorMessage)
	UpdateSocialMediaSvc(ctx context.Context, socialMedia SocialMedia) (result SocialMedia, errMsg message.ErrorMessage)
	DeleteSocialMediaSvc(ctx context.Context, socmedId uint64) (errMsg message.ErrorMessage)
}
