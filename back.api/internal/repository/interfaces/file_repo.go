package interfaces

import (
	"context"
	"io"
	"main/internal/domain"
)

type FileRepository interface {
	Save(ctx context.Context, file io.Reader, filePath string, fileType domain.FileType, userID int, fileName string) (int, error)
	Delete(ctx context.Context, fileID int) error
	GetByID(ctx context.Context, fileID int) (*domain.File, error)
}