package usecase

import (
	"context"
	"io"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/errors"
	repoInterface "main/internal/repository/interfaces"
	svcInterface "main/internal/service/interfaces"
	"main/internal/usecase/interfaces"
)

type fileUsecaseImpl struct {
	cfg      *config.StorageConfig
	fileSvc  svcInterface.FileService
	fileRepo repoInterface.FileRepository
}

func NewFileUsecase(
	fs svcInterface.FileService,
	fr repoInterface.FileRepository,
	cfg *config.StorageConfig,
) interfaces.FileUsecase {
	return &fileUsecaseImpl{
		fileSvc:  fs,
		fileRepo: fr,
		cfg:      cfg,
	}
}

func (uc *fileUsecaseImpl) UploadFile(ctx context.Context, file io.Reader, fileName string, fileType domain.FileType, userID int) (int, error) {

	// 1. 확장자 검증
	if !uc.fileSvc.IsAllowedExtension(fileType) {
		return 0, errors.ErrUnsupportedFile
	}

	// 2. 파일 저장 (UUID 기반 경로 반환)
	path, err := uc.fileSvc.SaveFile(ctx, file, fileName, fileType)
	if err != nil {
		return 0, err
	}

	// 3. DB에 메타데이터 저장 (userID, 원본 파일명 포함)
	id, err := uc.fileRepo.Save(ctx, file, path, fileType, userID, fileName)
	if err != nil {
		return 0, errors.ErrDatabase.Wrap(err)
	}

	return id, nil
}

func (uc *fileUsecaseImpl) DownloadFile(ctx context.Context, fileID int, userID int, role domain.RoleType) (io.ReadCloser, *domain.File, error) {
	// 1. DB에서 파일 메타데이터 조회
	file, err := uc.fileRepo.GetByID(ctx, fileID)
	if err != nil {
		return nil, nil, err
	}

	// 게시글에 첨부된 파일은 해당 게시글을 볼 수 있는 모든 사용자에게 공개된다(게시글 자체가
	// 별도 접근 제어 없이 공개되어 있으므로). 아직 게시글에 첨부되지 않은 파일은 업로더 본인/관리자만 접근 가능.
	if file.PostID == nil && file.UserID != userID && role != domain.AdminRole {
		return nil, nil, errors.ErrForbidden
	}

	// 2. 파일 스트림 열기
	reader, err := uc.fileSvc.DownloadFile(ctx, file.FilePath)
	if err != nil {
		return nil, nil, err
	}

	return reader, file, nil
}

func (uc *fileUsecaseImpl) DeleteFile(ctx context.Context, fileID int, userID int, role domain.RoleType) error {
	file, err := uc.fileRepo.GetByID(ctx, fileID)
	if err != nil {
		return err
	}
	if file.UserID != userID && role != domain.AdminRole {
		return errors.ErrForbidden
	}
	return uc.fileRepo.Delete(ctx, fileID)
}