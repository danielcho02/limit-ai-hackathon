package usecase

import (
	"context"
	"main/internal/domain"
	"main/internal/errors"
	repoInterfaces "main/internal/repository/interfaces"
	usecaseInterfaces "main/internal/usecase/interfaces"
)

type PostUseCaseImpl struct {
	postRepo repoInterfaces.PostRepository
	fileRepo repoInterfaces.FileRepository
}

func NewPostUseCase(pr repoInterfaces.PostRepository) usecaseInterfaces.PostUsecase {
	return &PostUseCaseImpl{
		postRepo: pr,
	}
}

func (p *PostUseCaseImpl) CreatePost(ctx context.Context, post *domain.Post, files []int) (*domain.Post, error) {
	
	if post.CategoryID < 0 || post.CategoryID >= domain.CategoryTypeMax {
		return nil, errors.ErrInvalidCategoryID
	}

	if err := p.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}

	if err := p.postRepo.AttachFiles(ctx, post.ID, files); err != nil {
		return nil, err
	}
	
	return p.postRepo.GetByID(ctx, post.ID)
}

func (p *PostUseCaseImpl) GetPosts(ctx context.Context, query domain.PostQuery) ([]domain.PostInfo, int, error) {
	
	if query.CategoryID != nil && (*query.CategoryID < 0 || *query.CategoryID >= int(domain.CategoryTypeMax)) {
		return nil, 0, errors.ErrInvalidCategoryID
	}

	total, err := p.postRepo.Count(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return []domain.PostInfo{}, 0, nil
	}

	postList, err := p.postRepo.GetMany(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	return postList, total, nil
}

// 상세 조회는 조회수를 1 올린 뒤 증가된 값을 그대로 내려준다.
// 수정/삭제 경로는 postRepo.GetByID를 직접 쓰므로 조회수가 오르지 않는다.
func (p *PostUseCaseImpl) GetPostDetail(ctx context.Context, postID int) (*domain.Post, error) {
	if err := p.postRepo.IncrementViews(ctx, postID); err != nil {
		return nil, err
	}
	return p.postRepo.GetByID(ctx, postID)
}

func (p *PostUseCaseImpl) UpdatePost(ctx context.Context, postID int, userID int, role domain.RoleType, title string, content string, categoryID domain.CategoryType, typeID domain.PostType, files []int) (*domain.Post, error) {
	post, err := p.postRepo.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}
	if post.AuthorID != userID && role != domain.AdminRole {
		return nil, errors.ErrForbidden
	}

	if err := p.postRepo.Update(ctx, postID, title, content, categoryID, typeID); err != nil {
		return nil, err
	}

	if err := p.postRepo.AttachFiles(ctx, postID, files); err != nil {
		return nil, err
	}

	return p.postRepo.GetByID(ctx, postID)
}

func (p *PostUseCaseImpl) DeletePost(ctx context.Context, postID int, userID int, role domain.RoleType) error {
	post, err := p.postRepo.GetByID(ctx, postID)
	if err != nil {
		return err
	}
	if post.AuthorID != userID && role != domain.AdminRole {
		return errors.ErrForbidden
	}

	if err := p.postRepo.Delete(ctx, postID); err != nil {
		return err
	}
	return nil
}