package v1

type RegisterRequest struct {
	ID  int    `json:"id"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type LoginRequest struct {
	StudentID int    `json:"student_id"`
	Password  string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}