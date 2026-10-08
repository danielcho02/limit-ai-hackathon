package interfaces

import (
	"context"
	"main/internal/domain"
)

type PostUsecase interface {
	CreatePost(ctx context.Context, post *domain.Post, files []int) (*domain.Post, error)
	GetPosts(ctx context.Context, query domain.PostQuery) ([]domain.PostInfo, int, error)
	GetPostDetail(ctx context.Context, postID int) (*domain.Post, error)
	UpdatePost(ctx context.Context, postID int, userID int, role domain.RoleType, title string, content string, categoryID domain.CategoryType, files []int) (*domain.Post, error)
	DeletePost(ctx context.Context, postID int, userID int, role domain.RoleType) error
}