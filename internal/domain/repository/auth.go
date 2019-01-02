package repository

import (
	"context"

	"vincent.com/auth/internal/domain/model"
)

//AuthRepository - User repo interface
type AuthRepository interface {
	NewToken(context.Context, *model.User) (token string, err error)
}
