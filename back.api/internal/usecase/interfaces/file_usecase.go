package interfaces

import (
	"context"
	"io"
	"main/internal/domain"
)

type FileUsecase interface {
	UploadFile(ctx context.Context, file io.Reader, fileName string, fileType domain.FileType, userID int) (int, error)
	DownloadFile(ctx context.Context, fileID int, userID int, role domain.RoleType) (io.ReadCloser, *domain.File, error)
	DeleteFile(ctx context.Context, fileID int, userID int, role domain.RoleType) error
}