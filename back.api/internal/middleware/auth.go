package middleware

import (
	"main/internal/auth"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// extractToken 은 Authorization 헤더에서만 토큰을 읽는다.
// 쿼리스트링(?token=...)으로 토큰을 전달하는 방식은 지원하지 않는다 — 로그/Referer/브라우저 히스토리에
// 토큰이 그대로 노출되는 문제가 있기 때문이다.
func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.ErrTokenMissing
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", errors.ErrInvalidAuthHeader
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.ErrInvalidAuthHeader
	}

	return token, nil
}

// appEnv: development 일 때만 web.front 의 개발용 목업 로그인 토큰("mock")을
// 허용한다. "mock" 은 실제 JWT가 아니라 development 에서만 유효하다.
func AuthMiddleware(
	cfg *config.AuthConfig,
	appEnv string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, errMessage := extractToken(c)
		if errMessage != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errMessage})
			c.Abort()
			return
		}

		if appEnv == "development" && tokenString == "mock" {
			c.Set("userID", 9999999999)
			c.Set("userRole", domain.AdminRole)
			c.Next()
			return
		}

		claims, err := auth.ValidateJWT(tokenString, []byte(cfg.AccessSecret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}