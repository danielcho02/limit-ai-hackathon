package interfaces

import (
	"context"
	"main/internal/domain"
)

//go:generate counterfeiter -o ../mocks --fake-name FakeUserUsecase . UserUsecase
type UserUsecase interface {
	LoginUser(ctx context.Context, studentID int, password string) (string, string, error)
	GetUserByID(ctx context.Context, userID int) (*domain.User, error)
	Refresh(ctx context.Context, refreshToken string) (string, error)
}