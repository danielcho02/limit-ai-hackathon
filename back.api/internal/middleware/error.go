package middleware

import (
	"main/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := errors.IsAppError(err); ok {
				if appErr.HTTPCode >= http.StatusInternalServerError {
				}

				c.JSON(appErr.HTTPCode,
				gin.H{
					"error": gin.H{
						"code":   appErr.Code,
						"message": appErr.Message,
					},
				})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"code":    errors.InternalServerErrorCode,
						"message": "Internal Server Error",
					},
				})
		}
 	}
}