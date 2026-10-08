package repository

import (
	"context"
	stdErrors "errors"
	"main/internal/domain"
	"main/internal/errors"
	"main/internal/repository/interfaces"

	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepoImpl{
		db: db,
	}
}

func (r *userRepoImpl) Create(ctx context.Context, user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.ErrUserAlreadyExists
		}
		return errors.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r *userRepoImpl) GetRoleByUserID(ctx context.Context, userID int) (domain.RoleType, error) {
	var user domain.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return 0, errors.ErrUserNotFound.Wrap(err)
	}
	return user.Role, nil
}

func (r *userRepoImpl) GetByID(ctx context.Context, userID int) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, errors.ErrUserNotFound.Wrap(err)
	}
	return &user, nil
}

func (r *userRepoImpl) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.ErrUserNotFound.Wrap(err)
	}
	return &user, nil
}