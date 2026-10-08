package presenter

import (
	"main/internal/domain"
	v1 "main/internal/handler/v1"
	"main/internal/utils"
)


func ToGetPostsResponse(postInfos[]domain.PostInfo, pagination utils.OffsetPagination, totalItems int) *v1.GetPostsResponse {
	pgMetadata := pagination.GetMetadata(totalItems)
	
	var postList []v1.PostInfo
	for _, info := range postInfos {
		postList = append(postList, v1.PostInfo{
			ID:           info.ID,
			Title:        info.Title,
			AuthorName:   info.AuthorName,
			CategoryID:   info.CategoryID,
			TypeID: info.TypeID,
			Views:        info.Views,
			CommentCount: info.CommentCount,
			CreatedAt:    formatTime(info.CreatedAt),
		})
	}

	return &v1.GetPostsResponse{
		Posts: postList,
		PaginationMetadata: pgMetadata,
	}
}

func PresentPost(post *domain.Post) *v1.GetPostDetailResponse {
	var files []v1.FileInfo
	for _, f := range post.Files {
		files = append(files, v1.FileInfo{
			ID:       f.ID,
			FileName: f.FileName,
		})
	}

	var comments []v1.CommentInfo
	for _, c := range post.Comments {
		comments = append(comments, v1.CommentInfo{
			ID:          c.ID,
			Content:     c.Content,
			AuthorID:    c.UserID,
			AuthorName:  c.AuthorNickname,
			CreatedAt:   formatTime(c.CreatedAt),
		})
	}

	return &v1.GetPostDetailResponse{
		ID:          post.ID,
		Title:       post.Title,
		AuthorID:    post.AuthorID,
		AuthorName:  post.AuthorName,
		CategoryID:  post.CategoryID,
		TypeID: post.TypeID,
		Content:     post.Content,
		Views:       post.Views,
		Files:       files,
		Comments:    comments,
		CreatedAt:   formatTime(post.CreatedAt),
		UpdatedAt:   formatTime(post.UpdatedAt),
	}
}

func PresentComment(c *domain.Comment) *v1.CommentInfo {
	return &v1.CommentInfo{
		ID:         c.ID,
		Content:    c.Content,
		AuthorID:   c.UserID,
		AuthorName: c.AuthorNickname,
		CreatedAt:  formatTime(c.CreatedAt),
	}
}