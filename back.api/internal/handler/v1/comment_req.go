package v1

type CreateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}