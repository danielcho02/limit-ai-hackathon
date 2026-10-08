package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"main/internal/auth"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/errors"
	"main/internal/repository/interfaces"
)

type authSvcImpl struct {
	userRepo interfaces.UserRepository
	authCfg  *config.AuthConfig
	smtpCfg *config.SMTPConfig
}

func NewAuthService(userRepo interfaces.UserRepository, authCfg *config.AuthConfig, smtpCfg *config.SMTPConfig) *authSvcImpl {
	return &authSvcImpl{
		userRepo: userRepo,
		authCfg:  authCfg,
		smtpCfg: smtpCfg,
	}
}

func (s *authSvcImpl) HashCode(ctx context.Context, code string, salt string) string {

	combined := code + salt

	// SHA-256 해시 생성
	hash := sha256.New()
	hash.Write([]byte(combined))
	
	// 해시 결과를 바이트 슬라이스로 받고, 이를 16진수 문자열로 변환
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

func (s *authSvcImpl) RegisterUser(ctx context.Context, ID int, email string, name string, enrollyear int, birthday string, password string) error {

	hashedPassword, err := auth.EncryptPassword(password)
	if err != nil {
		return errors.ErrInternalServer.Wrap(err)
	}

	user := &domain.User{
		ID:    ID,
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         domain.DefaultRole, // Set default role
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *authSvcImpl) GetUserRole(ctx context.Context, userID int) (domain.RoleType, error) {
	role, err := s.userRepo.GetRoleByUserID(ctx, userID)
	if err != nil {
		return domain.DefaultRole, err
	}
	return role, nil
}

func (s *authSvcImpl) IsAdmin(ctx context.Context, userID int) (bool, error) {
	role, err := s.userRepo.GetRoleByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	return role == domain.AdminRole, nil
}