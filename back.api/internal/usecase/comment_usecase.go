package usecase

import (
	"context"
	"main/internal/domain"
	"main/internal/errors"
	repoInterfaces "main/internal/repository/interfaces"
	"main/internal/usecase/interfaces"
)

type CommentUsecase struct {
	commentRepo repoInterfaces.CommentRepository
}

func NewCommentUsecase(commentRepo repoInterfaces.CommentRepository) interfaces.CommentUsecase {
	return &CommentUsecase{
		commentRepo: commentRepo,
	}
}

func (u *CommentUsecase) CreateComment(ctx context.Context, postID int, userID int, content string) (*domain.Comment, error) {
	comment := &domain.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}
	if err := u.commentRepo.Create(ctx, comment); err != nil {
		return nil, err
	}
	
	createdComment, err := u.commentRepo.GetByID(ctx, comment.ID)
	if err != nil {
		return nil, err
	}
	return createdComment, nil
}

func (u *CommentUsecase) GetCommentsByPostID(ctx context.Context, postID int) ([]*domain.Comment, error) {
	return u.commentRepo.GetCommentsByPostID(ctx, postID)
}

func (u *CommentUsecase) DeleteComment(ctx context.Context, commentID int, userID int, role domain.RoleType) error {
	comment, err := u.commentRepo.GetByID(ctx, commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID && role != domain.AdminRole {
		return errors.ErrForbidden
	}
	return u.commentRepo.Delete(ctx, commentID)
}