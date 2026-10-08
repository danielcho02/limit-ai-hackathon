package repository

import (
	"context"
	stdErrors "errors"
	"io"
	"main/internal/domain"
	"main/internal/errors"
	"main/internal/repository/interfaces"

	"gorm.io/gorm"
)

type fileRepoImpl struct {
	db      *gorm.DB
}

// NewFileeRepository 는 로컬 파일시스템 기반 StorageRepository 를 생성합니다.
// baseDir: 파일이 저장될 상대 경로 (예: "uploads")
func NewFileRepository(db *gorm.DB) interfaces.FileRepository {
	return &fileRepoImpl{
		db: db,
	}
}

func (r *fileRepoImpl) Save(ctx context.Context, file io.Reader, filePath string, fileType domain.FileType, userID int, fileName string) (int, error) {

	// DB에 메타데이터 저장
	record := &domain.File{
		UserID:    userID,
		FileName:  fileName,
		FilePath:  filePath,
		Extension: fileType,
	}
	if err := r.db.WithContext(ctx).Create(record).Error; err != nil {
		return 0, errors.ErrDatabase.Wrap(err)
	}

	return record.ID, nil
}


func (r *fileRepoImpl) Delete(ctx context.Context, fileID int) error {
	err := r.db.WithContext(ctx).Delete(&domain.File{}, fileID).Error
	if err != nil {
		return errors.ErrDatabase.Wrap(err)
	}
	return nil
}

func (r *fileRepoImpl) GetByID(ctx context.Context, fileID int) (*domain.File, error) {
	var file domain.File
	err := r.db.WithContext(ctx).First(&file, fileID).Error
	if err != nil {
		if stdErrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrFileNotFound.Wrap(err)
		}
		return nil, errors.ErrDatabase.Wrap(err)
	}
	return &file, nil
}