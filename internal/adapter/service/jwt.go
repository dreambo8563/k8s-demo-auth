package service

import (
	"context"
	"errors"

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

//ParseToken -
func (a *AuthRepository) ParseToken(ctx context.Context, t string) (u *model.User, err error) {
	uid, err := jwt.Parse(ctx, t)
	if err != nil {
		if jwt.IsExpired(err) {
			return nil, errors.New("token is Expired")
		}
		return nil, errors.New("token is Malformed")
	}
	return &model.User{
		ID: uid,
	}, nil
}
