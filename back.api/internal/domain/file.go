package domain

import (
	"time"
)

type FileType string

type File struct {
	ID       int    `gorm:"primaryKey"` // 파일 고유 pk

	// 외래키, nullable
	// logic 흐름 상 파일이 먼저 생성되고, 이후에 게시물에 첨부되는 형태이므로 nullable로 설정
	PostID   *int   `gorm:"index"`

	UserID   int    // 업로드한 사용자 ID
	FileName string // 원본 파일 이름
	FilePath string // 저장된 파일 경로
	Extension FileType // 파일 확장자 (예: .jpg, .pdf)

	CreatedAt time.Time
}

// mimeTypes maps file extensions to MIME types and inline rendering flag.
// inline=true: 브라우저에서 바로 열 수 있는 타입 (txt, html, 이미지 등)
// inline=false: 다운로드 처리 (pdf, mp4 등)
var mimeTypes = map[FileType]struct {
	MIME   string
	Inline bool
}{
	".jpg":  {"image/jpeg", true},
	".jpeg": {"image/jpeg", true},
	".png":  {"image/png", true},
	".gif":  {"image/gif", true},
	".txt":  {"text/plain; charset=utf-8", true},
	".html": {"text/html; charset=utf-8", true},
	".pdf":  {"application/pdf", false},
	".mp4":  {"video/mp4", false},
}

// MIMEType 은 파일 확장자에 해당하는 MIME 타입과 브라우저 inline 렌더링 여부를 반환합니다.
// 알 수 없는 확장자의 경우 application/octet-stream 으로 fallback 합니다.
func (f *File) MIMEType() (mime string, inline bool) {
	if info, ok := mimeTypes[f.Extension]; ok {
		return info.MIME, info.Inline
	}
	return "application/octet-stream", false
}