package usecase

import (
	"context"

	"main/internal/auth"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/errors"
	repository "main/internal/repository/interfaces"
	ucInterfaces "main/internal/usecase/interfaces"
)

type userUsecaseImpl struct {
	userRepo repository.UserRepository
	cfg 	*config.AuthConfig
}

func NewUserUsecase(ur repository.UserRepository, authCfg *config.AuthConfig) ucInterfaces.UserUsecase {
	return &userUsecaseImpl{
		userRepo: ur,
		cfg:      authCfg,
	}
}

func (s *userUsecaseImpl) LoginUser(ctx context.Context, ID int, password string) (string, string, error) {
	user, err := s.userRepo.GetByID(ctx, ID)
	if err != nil {
		return "", "", errors.ErrUserNotFound.Wrap(err)
	}
	if err := auth.ComparePassword(user.PasswordHash, password); err != nil {
		return "", "", errors.ErrPasswordMismatch
	}

	accessToken, err := auth.GenerateJWT(user.ID, user.Role, []byte(s.cfg.AccessSecret), s.cfg.AccessDuration)
	if err != nil {
		return "", "", errors.ErrInternalServer.Wrap(err)
	}
	refreshToken, err := auth.GenerateJWT(user.ID, user.Role, []byte(s.cfg.RefreshSecret), s.cfg.RefreshDuration)
	if err != nil {
		return "", "", errors.ErrInternalServer.Wrap(err)
	}
	return accessToken, refreshToken, nil
}

func (s *userUsecaseImpl) GetUserByID(ctx context.Context, userID int) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *userUsecaseImpl) Refresh(ctx context.Context, refreshToken string) (string, error) {
	claims, err := auth.ValidateJWT(refreshToken, []byte(s.cfg.RefreshSecret))
	if err != nil {
		return "", errors.ErrUnauthorized.Wrap(err)
	}
	
	role := claims.Role

	// Generate a new access token
	newAccessToken, err := auth.GenerateJWT(claims.UserID, role, []byte(s.cfg.AccessSecret), s.cfg.AccessDuration)
	if err != nil {
		return "", errors.ErrInternalServer.Wrap(err)
	}

	return newAccessToken, nil
}
