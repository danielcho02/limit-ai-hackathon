package v1

import "main/internal/domain"

type RegisterRequest struct {
	ID  int    `json:"id"`
	Nickname string `json:"nickname"`
	Password   string `json:"password"`
	Role domain.RoleType `json:"role"`
}

type LoginRequest struct {
	ID int    `json:"id"`
	Password  string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}