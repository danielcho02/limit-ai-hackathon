package interfaces

import (
	"context"
	"main/internal/domain"
)

type CommentUsecase interface {
	CreateComment(ctx context.Context, postID int, userID int, content string) (*domain.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID int) ([]*domain.Comment, error)
	DeleteComment(ctx context.Context, commentID int, userID int, role domain.RoleType) error
}