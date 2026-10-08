package service

import (
	"context"
	"fmt"
	"io"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/service/interfaces"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

var allowedExtensions = map[domain.FileType]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
	".pdf": true, ".mp4": true,
}

type fileServiceImpl struct {
	base string
}

func NewFileService(cfg *config.StorageConfig) interfaces.FileService {
	return &fileServiceImpl{
		base: cfg.UploadPath,
	}
}

func (s *fileServiceImpl) IsAllowedExtension(fileType domain.FileType) bool {
	return allowedExtensions[fileType]
}

// SaveFile 은 파일을 날짜/UUID 기반 경로에 저장하고, base 기준 상대 경로를 반환합니다.
// 동일 파일명 충돌을 방지하기 위해 UUID를 파일명으로 사용합니다.
// 반환 경로 예시: "2026/06/05/550e8400-e29b-41d4-a716-446655440000.png"
func (s *fileServiceImpl) SaveFile(ctx context.Context, file io.Reader, fileName string, fileType domain.FileType) (string, error) {
	now := time.Now()
	// UUID + 원본 확장자로 고유 파일명 생성
	uniqueName := uuid.New().String() + string(fileType)
	// 날짜 기반 디렉토리로 분산 저장
	relPath := filepath.Join(
		fmt.Sprintf("%d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%02d", now.Day()),
		uniqueName,
	)
	absPath := filepath.Join(s.base, relPath)

	// 중간 디렉토리 생성
	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return "", fmt.Errorf("storage: mkdir %q: %w", filepath.Dir(absPath), err)
	}

	// 파일 기록
	dst, err := os.Create(absPath)
	if err != nil {
		return "", fmt.Errorf("storage: create file %q: %w", absPath, err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("storage: write file %q: %w", absPath, err)
	}

	// base 기준 상대 경로 반환 (DB 저장용)
	return filepath.ToSlash(relPath), nil
}

func (s *fileServiceImpl) DownloadFile(ctx context.Context, filePath string) (io.ReadCloser, error) {
	absPath := filepath.Join(s.base, filepath.FromSlash(filePath))

	f, err := os.Open(absPath)
	if err != nil {
		return nil, fmt.Errorf("storage: open file %q: %w", absPath, err)
	}

	return f, nil
}