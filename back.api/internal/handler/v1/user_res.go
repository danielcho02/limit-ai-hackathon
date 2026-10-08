package v1

import "main/internal/domain"

type GetUserDetailResponse struct {
	ID int `json:"id"`
	Nickname string 	`json:"nickname"`
	Role     domain.RoleType `json:"role" gorm:"not null"`
	ProfileImageID int `json:"profile_image_id" gorm:"not null"`
}