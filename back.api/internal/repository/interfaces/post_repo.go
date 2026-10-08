package interfaces

import (
	"context"
	"main/internal/domain"
)

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetByID(ctx context.Context, id int) (*domain.Post, error)
	GetMany(ctx context.Context, query domain.PostQuery) ([]domain.PostInfo, error)
	Count(ctx context.Context, query domain.PostQuery) (int, error)

	AttachFiles(ctx context.Context, postID int, fileIDs []int) error
	IncrementViews(ctx context.Context, postID int) error
	Update(ctx context.Context, postID int, title string, content string, categoryID domain.CategoryType) error
	Delete(ctx context.Context, postID int) error
}

