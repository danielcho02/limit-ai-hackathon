package handler

import (
	"fmt"
	"main/internal/domain"
	"main/internal/errors"
	v1 "main/internal/handler/v1"
	repoInterface "main/internal/repository/interfaces"
	svcInterface "main/internal/service/interfaces"
	ucInterface "main/internal/usecase/interfaces"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	filesvc svcInterface.FileService
	fileUsecase ucInterface.FileUsecase
	repo repoInterface.FileRepository
}

func NewFileHandler(svc svcInterface.FileService, uc ucInterface.FileUsecase, repo repoInterface.FileRepository) *FileHandler {
	return &FileHandler{
		filesvc: svc,
		fileUsecase: uc,
		repo: repo,
	}
}

// UploadFile godoc
// @Summary      파일 업로드
// @Description  허용 확장자: .jpg, .jpeg, .png, .gif, .pdf, .mp4
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Tags         File
// @Produce      json
// @Param        files  formData  file  true  "업로드할 파일(single)"
// @Success      200    {object}  v1.Response{data=v1.UploadFileResponse}
// @Failure      400    {object}  errors.AppError  "ErrShouldBindJson(5001): multipart form 파싱 실패. / ErrBadRequest(5001): files 파트가 비어 있습니다. / ErrUnsupportedFile(5103): 지원하지 않는 파일 확장자입니다."
// @Failure      401    {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      500    {object}  errors.AppError  "ErrFileOpen(5102): 업로드 파일 열기 실패. / ErrDatabase(9999): 파일 메타데이터 저장 실패."
// @Router       /api/v1/files [post]
func (h *FileHandler) UploadFile(c *gin.Context) {
	ctx := c.Request.Context()

	userID, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.Error(errors.ErrBadRequest)
		return
	}

	// 단일 파일 업로드
	fileHeader := files[0]
	file, err := fileHeader.Open()
	if err != nil {
		c.Error(errors.ErrFileOpen.Wrap(err))
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	id, err := h.fileUsecase.UploadFile(ctx, file, fileHeader.Filename, domain.FileType(ext), userID.(int))
	if err != nil {
		c.Error(err)
		return
	}
	
	c.JSON(200, v1.Response{
		Message: "파일 업로드 성공",
		Data: v1.UploadFileResponse{
			FileID: id,
		},
	})
}

// DownloadFile godoc
// @Summary      파일 다운로드
// @Description  이미지, 텍스트, HTML 파일은 브라우저에서 직접 열리며, 나머지(webp, MP4 등)는 다운로드됩니다. 이미 게시글에 첨부된 파일은 누구나(인증된 사용자라면) 다운로드할 수 있고, 아직 게시글에 첨부되지 않은 파일은 업로더 본인 또는 Admin만 다운로드할 수 있습니다.
// @Security     BearerAuth
// @Tags         File
// @Param        file_id  path     int  true  "파일 ID"
// @Success      200      {file}  file
// @Failure      400      {object}  errors.AppError  "ErrInvalidFileID(5102): file_id가 올바른 정수 형식이 아닙니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      403      {object}  errors.AppError  "ErrForbidden(108): 게시글에 첨부되지 않은 파일에 대해 업로더 본인/Admin이 아닙니다."
// @Failure      404      {object}  errors.AppError  "ErrFileNotFound(5104): 존재하지 않는 파일입니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999) 등: 파일 스트림을 여는 중 내부 서버 오류"
// @Router       /api/v1/files/{file_id} [get]
func (h *FileHandler) DownloadFile(c *gin.Context) {
	ctx := c.Request.Context()

	strFileID := c.Param("file_id")
	fileID, err := strconv.Atoi(strFileID)
	if err != nil {
		c.Error(errors.ErrInvalidFileID.Wrap(err))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	role, _ := c.Get("userRole")
	userRole, _ := role.(domain.RoleType)

	reader, fileMeta, err := h.fileUsecase.DownloadFile(ctx, fileID, userID.(int), userRole)
	if err != nil {
		c.Error(err)
		return
	}
	defer reader.Close()

	mimeType, inline := fileMeta.MIMEType()

	// inline=true: 브라우저에서 직접 렌더링 (이미지, 텍스트, HTML)
	// inline=false: 다운로드 강제 (PDF, MP4 등)
	disposition := "attachment"
	if inline {
		disposition = "inline"
	}

	// RFC 5987에 따라 UTF-8 파일명 인코딩
	encodedName := url.PathEscape(fileMeta.FileName)
	c.Header("Content-Disposition", fmt.Sprintf(`%s; filename*=UTF-8''%s`, disposition, encodedName))
	c.DataFromReader(200, -1, mimeType, reader, nil)
}


// DeleteFile godoc
// @Summary      파일 삭제
// @Description  업로더 본인 또는 Admin만 삭제할 수 있습니다.
// @Security     BearerAuth
// @Tags         File
// @Produce      json
// @Param        file_id  path     int  true  "삭제할 파일 ID"
// @Success      204      {object}  v1.Response
// @Failure      400      {object}  errors.AppError  "ErrInvalidFileID(5102): file_id가 올바른 정수 형식이 아닙니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      403      {object}  errors.AppError  "ErrForbidden(108): 업로더 본인 또는 Admin이 아닙니다."
// @Failure      404      {object}  errors.AppError  "ErrFileNotFound(5104): 존재하지 않는 파일입니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999): 파일 삭제 실패"
// @Router       /api/v1/files/{file_id} [delete]
func (h *FileHandler) DeleteFile(c *gin.Context) {
	ctx := c.Request.Context()

	strFileID := c.Param("file_id")
	fileID, err := strconv.Atoi(strFileID)
	if err != nil {
		c.Error(errors.ErrInvalidFileID.Wrap(err))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	role, _ := c.Get("userRole")
	userRole, _ := role.(domain.RoleType)

	if err := h.fileUsecase.DeleteFile(ctx, fileID, userID.(int), userRole); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, v1.Response{
		Message: "파일 삭제 성공",
	})
}