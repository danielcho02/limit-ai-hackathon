package interfaces

import (
	"context"
	"main/internal/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	GetByID(ctx context.Context, commentID int) (*domain.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID int) ([]*domain.Comment, error)
	Delete(ctx context.Context, commentID int) error
}