package v1

type UpdateUserRequest struct {
	Nickname string `json:"nickname"`
	ProfileImageID string `json:"profile_image_id"`
}