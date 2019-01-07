package usecase

import (
	"context"

	"vincent.com/auth/internal/domain/model"

	"vincent.com/auth/internal/domain/repository"
)

// User -
type User struct {
	ID string `json:"uid"`
}

//AuthUsecase -
type AuthUsecase struct {
	repo repository.AuthRepository
}

//NewAuthUsecase -
func NewAuthUsecase(repo repository.AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

//NewToken -
func (a *AuthUsecase) NewToken(ctx context.Context, u *User) (token string, err error) {

	token, err = a.repo.NewToken(ctx, toUser(u))
	if err != nil {
		return "", err
	}

	return token, nil
}

//ParseToken -
func (a *AuthUsecase) ParseToken(ctx context.Context, t string) (u *User, err error) {

	user, err := a.repo.ParseToken(ctx, t)
	if err != nil {
		return nil, err
	}

	return fromUser(user), nil
}

func toUser(u *User) *model.User {
	return &model.User{
		ID: u.ID,
	}
}

func fromUser(u *model.User) *User {
	return &User{
		ID: u.ID,
	}
}
