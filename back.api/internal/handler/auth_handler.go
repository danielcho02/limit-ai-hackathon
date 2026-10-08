package handler

import (
	"main/internal/errors"
	v1 "main/internal/handler/v1"
	svcInterface "main/internal/service/interfaces"
	ucInterface "main/internal/usecase/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService svcInterface.AuthService
	userUsecase ucInterface.UserUsecase
}

func NewAuthHandler(authService svcInterface.AuthService, userUsecase ucInterface.UserUsecase) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userUsecase: userUsecase,
	}
}

// Register godoc
// @Summary      회원가입
// @Description  회원가입을 위한 API. 학번, 비밀번호, 이름, 이메일, 입학연도 정보를 받아 회원을 등록합니다. email_verification_id는 반드시 verify-email로 검증에 성공한(사용됨 처리된) 레코드여야 하며, 그 레코드의 email과 요청의 email이 일치해야 합니다.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      v1.RegisterRequest  true  "회원가입 요청 파라미터"
// @Success      201      {object}  v1.Response
// @Failure      400      {object}  errors.AppError  "ErrShouldBindJson(5001): 잘못된 JSON 형식입니다. / ErrEmailVerification(5201): 이메일 인증이 완료되지 않았습니다. / ErrEmailVerificationExpired(5201): 이메일 인증이 만료되었습니다. / ErrInvalidResetPasswordRequest(5201): 인증된 이메일과 요청 이메일이 일치하지 않습니다. / ErrUserAlreadyExists(103): 이미 존재하는 사용자입니다."
// @Failure      404      {object}  errors.AppError  "ErrEmailVerificationNotFound(5202): email_verification_id에 해당하는 인증 요청을 찾을 수 없습니다."
// @Failure      500      {object}  errors.AppError  "ErrInternalServer(5000): 비밀번호 암호화 실패 / ErrDatabase(9999): 사용자 저장 실패"
// @Router       /api/v1/auth/register [post]
func (h *AuthHandler) RegisterLocalUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req v1.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	err := h.authService.RegisterUser(ctx, req.ID, req.Email, req.Name, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	// 성공 시 201 Created
	c.JSON(http.StatusCreated, v1.Response{
		Message: "Success",
		Data:    nil,
	})
}

// Login godoc
// @Summary      로그인
// @Description  로그인 및 토큰 발급 API
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      v1.LoginRequest  true  "로그인 요청 파라미터"
// @Success      200      {object}  v1.Response{data=v1.LoginResponse}
// @Failure      400  {object}  errors.AppError  "ErrShouldBindJson(5001): 잘못된 JSON 형식입니다."
// @Failure      401  {object}  errors.AppError  "ErrPasswordMismatch(100): 비밀번호가 일치하지 않습니다."
// @Failure      404  {object}  errors.AppError  "ErrUserNotFound(101): 존재하지 않는 학번입니다."
// @Failure      500  {object}  errors.AppError  "ErrInternalServer(5000): 토큰 발급 실패"
// @Router       /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req v1.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	accessToken, refreshToken, err := h.userUsecase.LoginUser(ctx, req.StudentID, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, v1.Response{
		Message: "Success",
		Data: v1.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})
}

// RefreshToken godoc
// @Summary      토큰 갱신
// @Description  Refresh Token을 사용하여 새로운 Access Token 발급
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      v1.RefreshRequest  true  "토큰 갱신 파라미터"
// @Success      200      {object}  v1.Response{data=v1.AuthRefreshResponse}
// @Failure      400      {object}  errors.AppError  "ErrShouldBindJson(5001): 잘못된 JSON 형식입니다."
// @Failure      401      {object}  errors.AppError  "ErrUnauthorized(106): 유효하지 않거나 만료된 refresh token입니다."
// @Failure      500      {object}  errors.AppError  "ErrInternalServer(5000): 토큰 발급 실패"
// @Router       /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()

	var req v1.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	newAccessToken, err := h.userUsecase.Refresh(ctx, req.RefreshToken)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, v1.Response{
		Message: "Success",
		Data: v1.AuthRefreshResponse{
			AccessToken: newAccessToken,
		},
	})
}