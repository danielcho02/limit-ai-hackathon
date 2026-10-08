package v1

import (
	"main/internal/domain"
	"main/internal/utils"
)

type PostInfo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	AuthorName string `json:"author_name"`
	CategoryID domain.CategoryType `json:"category_id"`
	TypeID domain.PostType `json:"type_id"`
	Views int `json:"views"`
	CommentCount int `json:"comment_count"`
	CreatedAt string `json:"created_at"`
}

type CreatePostResponse struct {
	Title string `json:"title" validate:"required"`

	Content string `json:"content" validate:"required"`

	CategoryID int `json:"category_id" validate:"required"`

	TypeID domain.PostType `json:"type_id"`

	Files []int `json:"files"`
}

type GetPostsResponse struct {
	Posts []PostInfo `json:"posts"`
	utils.PaginationMetadata `json:"pagination"`
}

type FileInfo struct {
	ID int `json:"id"`
	FileName string `json:"file_name"`
}

type CommentInfo struct {
	ID int `json:"id"`
	Content string `json:"content"`
	AuthorID int `json:"author_id"`
	AuthorName string `json:"author_name"`
	CreatedAt string `json:"created_at"`
}

type GetPostDetailResponse struct {
	ID int `json:"id"`

	Title string `json:"title"`

	Content string `json:"content"`

	AuthorID int `json:"author_id"`
	AuthorName string `json:"author_name"`
	CategoryID domain.CategoryType `json:"category_id"`
	TypeID domain.PostType `json:"type_id"`

	Views int `json:"views"`

	Files []FileInfo `json:"files,omitempty"`
	Comments []CommentInfo `json:"comments,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
