package service

import (
	"context"

	"vincent.com/auth/internal/pkg/jwt"

	"vincent.com/auth/internal/domain/model"

	"vincent.com/auth/internal/pkg/logger"
)

var log = logger.Logger()

//AuthRepository -
type AuthRepository struct {
}

//NewAuthRepository -
func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

//NewToken -
func (a *AuthRepository) NewToken(ctx context.Context, u *model.User) (token string, err error) {
	return jwt.New(ctx, u.GetID())
}
