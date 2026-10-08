package v1

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthValidResponse struct {
	Token string `json:"token"`
}

type SigninResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}