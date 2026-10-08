package main

import (
	"main/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func wireUserDependency(r *gin.Engine, authMiddleware gin.HandlerFunc, db *gorm.DB, cfg *config.Config) {
	
}