package handler

import (
	"main/internal/domain"
	"main/internal/errors"
	v1 "main/internal/handler/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler{
	return &UserHandler{
		db : db,
	}
}

// @UpdateProfile godoc
// @Summary      사용자 프로필 정보 수정
// @Description  사용자가 자신의 닉네임과 프로필 이미지를 수정할 수 있는 엔드포인트입니다.
// @Security     BearerAuth
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      v1.UpdateUserRequest  true  "유저 정보 수정 요청 파라미터"
// @Success      200 {object} v1.Response
// @Failure      400 {object} errors.AppError "ErrShouldBindJson(5001): 잘못된 JSON 형식입니다. / ErrBadRequest(5001): nickname 또는 profile_image_id가 누락되었습니다."
// @Failure      401 {object} errors.AppError "ErrTokenMissing(105): 인증 토큰이 누락되었습니다. / ErrInvalidAuthHeader(104): 유효하지 않은 Authorization 헤더입니다."
// @Failure      403 {object} errors.AppError "ErrUnauthorized(106): 인증되지 않은 사용자입니다."
// @Failure      500 {object} errors.AppError "ErrDatabase(9999): 데이터베이스 처리 중 오류가 발생했습니다."
// @Router       /api/v1/user/profile [patch]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	ctx := c.Request.Context()

	userID, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}

	var req v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	// update logic
	updateData := map[string]interface{}{
		"nickname": req.Nickname,
		"profile_image_id": req.ProfileImageID,
	}

	result := h.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", userID).
		Updates(updateData)

	if result.Error != nil {
		c.Error(errors.ErrDatabase.Wrap(result.Error))
		return
	}

	if result.RowsAffected == 0 {
		c.Error(errors.ErrPostNotFound)
		return
	}

	
	c.JSON(200, v1.Response{
		Message: "프로필 수정 성공",
	})

}