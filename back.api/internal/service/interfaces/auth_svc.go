package interfaces

import (
	"context"
	"main/internal/domain"
)

//go:generate counterfeiter -o ../mocks --fake-name FakeAuthService . AuthService
type AuthService interface {
	HashCode(ctx context.Context, code string, salt string) string

	RegisterUser(ctx context.Context, ID int, email string, name string, password string) error
	IsAdmin(ctx context.Context, userID int) (bool, error)

	GetUserRole(ctx context.Context, userID int) (domain.RoleType, error)
}