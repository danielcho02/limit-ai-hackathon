package v1

import "main/internal/domain"

type CreatePostRequest struct {
	Title   string `json:"title"`
	Files []int `json:"files"`
	Content string `json:"content"`
	CategoryID domain.CategoryType `json:"category_id"`
	TypeID domain.PostType `json:"type_id"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Files []int `json:"files"`
	Content string `json:"content"`
	CategoryID domain.CategoryType `json:"category_id"`
	TypeID domain.PostType `json:"type_id"`
}