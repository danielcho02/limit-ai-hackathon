package domain

import "time"

type Comment struct {
	ID        int      `gorm:"primaryKey" json:"id"`

	PostID    int      `gorm:"not null" json:"post_id"`
	UserID    int      `gorm:"not null" json:"user_id"`
	AuthorNickname string `gorm:"->" json:"author_nickname"`

	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentInfo struct {
	UserName  string    `json:"user_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}