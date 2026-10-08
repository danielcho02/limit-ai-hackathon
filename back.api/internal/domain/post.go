package domain

import (
	"time"
)

type CategoryType int

const (
	NoticePost CategoryType = iota // 0 공지글
	ArtPost // 1 예술
	EnvironmentPost // 2 환경
	EducationPost // 3 교육
	ConsultPost // 4 상담
	SportsPost // 5 스포츠
	BeautyPost // 6 미용
	ETCPost // 7 기타

	CategoryTypeMax
)

type PostType int

const (
	ProviderPost PostType = iota + 101
	RequestPost // 102
)

type Post struct {
	ID      int `gorm:"primaryKey"`

	Title  string
	
	AuthorID  int
	AuthorName string `gorm:"->"`
	CategoryID CategoryType `gorm:"index"`
	TypeID PostType `gorm:"index"`

	Content string

	Views int `gorm:"not null;default:0"`

	// 글 삭제 시 첨부파일/댓글도 함께 삭제되어야 한다(FK NO ACTION이면 삭제가 500으로 실패함).
	Files []File `gorm:"constraint:OnDelete:CASCADE;"`
	Comments []Comment `gorm:"constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
	DeletedAt *time.Time
}

type PostInfo struct {
	ID int `json:"id"`

	Title string `json:"title"`
	AuthorName string `json:"author_name"`
	CategoryID CategoryType `json:"category_id"`
	TypeID PostType `json:"post_type"`
	Views int `json:"views"`
	CommentCount int `json:"comment_count"`
	CreatedAt time.Time `json:"created_at"`
}

type PostQuery struct {
	// options
	Title *string
	CategoryID *int
	TypeID *int
	AuthorID *int

	// offset pagination
	Limit  int
	Offset int
}
