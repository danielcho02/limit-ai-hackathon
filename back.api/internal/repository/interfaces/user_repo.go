package interfaces

import (
	"context"
	"main/internal/domain"
)

//go:generate counterfeiter -o ../mocks --fake-name FakeUserRepository . UserRepository
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetRoleByUserID(ctx context.Context, userID int) (domain.RoleType, error)
	GetByID(ctx context.Context, userID int) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}