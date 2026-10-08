package interfaces

import (
	"context"
	"io"
	"main/internal/domain"
)

type FileService interface {
	IsAllowedExtension(fileType domain.FileType) bool
	SaveFile(ctx context.Context, file io.Reader, fileName string, fileType domain.FileType) (string, error)
	DownloadFile(ctx context.Context, filePath string) (io.ReadCloser, error)
}