package v1

type RegisterRequest struct {
	ID  int    `json:"id"`
	Nickname string `json:"nickname"`
	Password   string `json:"password"`
	Name       string `json:"name"`
}

type LoginRequest struct {
	ID int    `json:"id"`
	Password  string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}