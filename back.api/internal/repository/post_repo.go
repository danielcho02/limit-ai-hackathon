package repository

import (
	"context"
	stdErrors "errors"
	"main/internal/domain"
	"main/internal/errors"
	"main/internal/repository/interfaces"

	"gorm.io/gorm"
)


type postRepoImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) interfaces.PostRepository {
	return &postRepoImpl{
		db: db,
	}
}

// PostQuery에 따라 쿼리를 build하는 헬퍼 함수
func buildQuery(
	ctx context.Context,
    db *gorm.DB,
	query domain.PostQuery,
) *gorm.DB {

	db = db.WithContext(ctx).
		Model(&domain.Post{}).
		Joins("LEFT JOIN users ON users.id = posts.author_id")

	if query.Title != nil && *query.Title != "" {
		db = db.Where(
			"posts.title LIKE ?",
			"%"+*query.Title+"%",
		)
	}

	if query.CategoryID != nil {
		db = db.Where(
			"posts.category_id = ?",
			*query.CategoryID,
		)
	}

	if query.AuthorID != nil {
		db = db.Where(
			"posts.author_id = ?",
			*query.AuthorID,
		)
	}

	return db
}

func (r *postRepoImpl) Create(
    ctx context.Context,
    post *domain.Post,
) error {
    if err := r.db.WithContext(ctx).
        Create(post).Error; err != nil {
        return errors.ErrDatabase.Wrap(err)
    }
    return nil
}

func (r *postRepoImpl) GetByID(ctx context.Context, id int) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.WithContext(ctx).
		Model(&domain.Post{}).
		Select("posts.*, users.nickname AS author_name").
		Joins("LEFT JOIN users ON users.id = posts.author_id").
		Preload("Files").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Select("comments.*, users.nickname AS author_nickname").
				Joins("LEFT JOIN users ON users.id = comments.user_id").
				Order("comments.created_at ASC")
		}).
		First(&post, id).Error; err != nil {
            if stdErrors.Is(err, gorm.ErrRecordNotFound) {
                return nil, errors.ErrPostNotFound.Wrap(err)
            }
        return nil, errors.ErrDatabase.Wrap(err) // 404 외 DB 에러
    }
    return &post, nil
}

// PostInfo 리스트 반환
// GetMany에서는 file, content 등을 포함하지 않음
func (r *postRepoImpl) GetMany(ctx context.Context, query domain.PostQuery) ([]domain.PostInfo, error) {
    var postList []domain.PostInfo

	db := buildQuery(ctx,r.db, query)

    if err := db.Select("posts.id, posts.title, posts.category_id, posts.views, posts.created_at, users.nickname as author_name, (SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) as comment_count").
        Scopes(paginate(query.Offset, query.Limit)).
        Order("posts.created_at DESC").
        Scan(&postList).Error; err != nil {
			if stdErrors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.ErrPostNotFound.Wrap(err)
			}
        return nil, errors.ErrDatabase.Wrap(err)
    }

	return postList, nil

}

func (r *postRepoImpl) Count(ctx context.Context, query domain.PostQuery) (int, error) {
	var total int64

	db := buildQuery(ctx, r.db, query)

    if err := db.Count(&total).Error; err != nil {
        return 0, errors.ErrDatabase.Wrap(err)
    }

	return int(total), nil
}

func (r *postRepoImpl) AttachFiles(
    ctx context.Context,
    postID int,
    fileIDs []int,
) error {

    var post domain.Post

    if err := r.db.WithContext(ctx).
        First(&post, postID).Error; err != nil {

        return errors.ErrPostNotFound.Wrap(err)
    }

    if len(fileIDs) == 0 {
        return nil
    }

    var files []domain.File

    if err := r.db.WithContext(ctx).
        Find(&files, fileIDs).Error; err != nil {

        return errors.ErrDatabase.Wrap(err)
    }

    if len(files) != len(fileIDs) {
        return errors.ErrFileNotFound
    }

    if err := r.db.WithContext(ctx).
        Model(&post).
        Association("Files").
        Replace(files); err != nil {

        return errors.ErrDatabase.Wrap(err)
    }

    return nil
}

func (r *postRepoImpl) IncrementViews(ctx context.Context, postID int) error {
	result := r.db.WithContext(ctx).
		Model(&domain.Post{}).
		Where("id = ?", postID).
		UpdateColumn("views", gorm.Expr("views + 1"))

	if result.Error != nil {
		return errors.ErrDatabase.Wrap(result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.ErrPostNotFound
	}

	return nil
}

func (r *postRepoImpl) Update(
	ctx context.Context,
	postID int,
	title string,
	content string,
	categoryID domain.CategoryType,
	typeID domain.PostType,
) error {

	updateData := map[string]interface{}{
		"title":       title,
		"content":     content,
		"category_id": categoryID,
		"type_id": typeID,
	}

	result := r.db.WithContext(ctx).
		Model(&domain.Post{}).
		Where("id = ?", postID).
		Updates(updateData)

	if result.Error != nil {
		return errors.ErrDatabase.Wrap(result.Error)
	}

	if result.RowsAffected == 0 {
		return errors.ErrPostNotFound
	}

	return nil
}

func (r *postRepoImpl) Delete(ctx context.Context, postID int) error {
	if err := r.db.Delete(&domain.Post{}, postID).Error; err != nil {
        if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrPostNotFound.Wrap(err)
		}
		return errors.ErrDatabase.Wrap(err)
    }
    return nil
}