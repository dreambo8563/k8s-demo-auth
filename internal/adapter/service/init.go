package service

import (
	"vincent.com/auth/internal/domain/usecase"
)

//InitializeAuthCase -
func InitializeAuthCase() *usecase.AuthUsecase {
	repo := NewAuthRepository()
	return usecase.NewAuthUsecase(repo)
}
