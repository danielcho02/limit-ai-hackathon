package repository

import (
	"context"
	stdErrors "errors"
	"main/internal/domain"
	"main/internal/errors"
	"main/internal/repository/interfaces"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) interfaces.CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment *domain.Comment) error {
	if err := r.db.WithContext(ctx).Create(comment).Error; err != nil {
		return errors.ErrDatabase.Wrap(err)
	}

	return nil
}

func (r *CommentRepository) GetByID(ctx context.Context, commentID int) (*domain.Comment, error) {
	var comment domain.Comment

	err := r.db.WithContext(ctx).
		Model(&domain.Comment{}).
		Select("comments.*, users.name AS author_nickname").
		Joins("LEFT JOIN users ON users.id = comments.user_id").
		Where("comments.id = ?", commentID).
		First(&comment).Error

	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrCommentNotFound.Wrap(err)
		}
		return nil, errors.ErrDatabase.Wrap(err)
	}

	return &comment, nil
}

func (r *CommentRepository) GetCommentsByPostID(ctx context.Context, postID int) ([]*domain.Comment, error) {
	var comments []*domain.Comment

	err := r.db.WithContext(ctx).
		Model(&domain.Comment{}).
		Select("comments.id, comments.content, comments.created_at, users.name AS author_nickname").
		Joins("JOIN users ON users.id = comments.user_id").
		Where("comments.post_id = ?", postID).
		Order("comments.created_at ASC").
		Scan(&comments).Error

	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrCommentNotFound.Wrap(err)
		}
		return nil, errors.ErrDatabase.Wrap(err)
	}

	return comments, nil
}

func (r *CommentRepository) Delete(ctx context.Context, commentID int) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Comment{}, commentID).Error; err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrCommentNotFound.Wrap(err)
		}
		return errors.ErrDatabase.Wrap(err)
	}
	return nil
}