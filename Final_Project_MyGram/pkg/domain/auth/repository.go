package auth

import (
	"context"

	"github.com/muhammadyubi/GoExcercise/tree/main/Final_Project_MyGram/pkg/domain/user"
)

type AuthRepo interface {
	LoginUser(ctx context.Context, username string) (result user.User, err error)
}
