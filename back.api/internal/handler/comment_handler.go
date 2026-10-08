package handler

import (
	"strconv"

	"main/internal/domain"
	"main/internal/errors"
	v1 "main/internal/handler/v1"
	"main/internal/presenter"
	"main/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentUsecase interfaces.CommentUsecase
}

func NewCommentHandler(commentUsecase interfaces.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		commentUsecase: commentUsecase,
	}
}

// CreateComment godoc
// @Summary      댓글 생성
// @Security     BearerAuth
// @Tags         Comment
// @Accept       json
// @Produce      json
// @Param        post_id  path      int                      true  "게시글 ID"
// @Param        comment  body      v1.CreateCommentRequest  true  "댓글 생성 요청"
// @Success      201      {object}  v1.Response{data=v1.CommentInfo}
// @Failure      400      {object}  errors.AppError  "ErrInvalidPostID(1002): post_id가 올바른 정수 형식이 아닙니다. / ErrShouldBindJson(5001): 잘못된 JSON 형식입니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      404      {object}  errors.AppError  "(참고) 현재 구현은 post_id 존재 여부를 사전 검증하지 않아 실제로는 반환되지 않습니다 - 존재하지 않는 post_id는 500(ErrDatabase)으로 처리됩니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999): 댓글 저장 실패 / 존재하지 않는 post_id로 인한 외래키 위반"
// @Router       /api/v1/posts/{post_id}/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	postIDStr := c.Param("post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.Error(errors.ErrInvalidPostID.Wrap(err))
		return
	}

	userIDVal, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		c.Error(errors.ErrUnauthorized)
		return
	}

	var req v1.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	comment, err := h.commentUsecase.CreateComment(ctx, postID, userID, req.Content)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, v1.Response{
		Message: "댓글 생성 성공",
		Data:    presenter.PresentComment(comment),
	})
}

// DeleteComment godoc
// @Summary      댓글 삭제
// @Description  작성자 본인 또는 Admin만 삭제할 수 있습니다.
// @Security     BearerAuth
// @Tags         Comment
// @Produce      json
// @Param        comment_id  path      int  true  "삭제할 댓글 ID"
// @Success      204         {object}  v1.Response
// @Failure      400         {object}  errors.AppError  "ErrInvalidCommentID(4002): comment_id가 올바른 정수 형식이 아닙니다."
// @Failure      401         {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      403         {object}  errors.AppError  "ErrForbidden(108): 작성자 본인 또는 Admin이 아닙니다."
// @Failure      404         {object}  errors.AppError  "ErrCommentNotFound(4001): 존재하지 않는 댓글입니다."
// @Failure      500         {object}  errors.AppError  "ErrDatabase(9999): 댓글 삭제 실패"
// @Router       /api/v1/comments/{comment_id} [delete]
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	ctx := c.Request.Context()

	commentIDStr := c.Param("comment_id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		c.Error(errors.ErrInvalidCommentID.Wrap(err))
		return
	}

	userIDVal, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	userID, ok := userIDVal.(int)
	if !ok {
		c.Error(errors.ErrUnauthorized)
		return
	}
	role, _ := c.Get("userRole")
	userRole, _ := role.(domain.RoleType)

	if err := h.commentUsecase.DeleteComment(ctx, commentID, userID, userRole); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, v1.Response{
		Message: "댓글 삭제 성공",
	})
}