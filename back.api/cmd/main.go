package main

import (
	"log"

	"main/internal/config"
	"main/internal/infra/db"
	"main/internal/middleware"

	database "main/internal/infra/db"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @title           title
// @version         1.0
// @description     description
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                "Bearer {access_token}" 형식으로 전달합니다. 로그인(`/auth/login`) 또는 회원가입(`/auth/register`) 응답으로 발급된 access_token을 사용하세요.
func main() {
	// 1. 설정 로드
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// ctx := context.Background()

	// 2. DB 연결 (GORM)
	db, err := db.InitPostgreSQL(*cfg.DB)
	if err != nil {
		log.Printf("Failed to connect to database : %v", err)
		panic(err)
	}

	// 3. AutoMigrate
	database.AutoMigrate(db)

	// 4. Gin 설정
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}


	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorMiddleware())

	authMiddleware := middleware.AuthMiddleware(cfg.Auth, cfg.App.Env)

	wireUserDependency(r, authMiddleware, db, cfg)

	r.Run(":8080")
}