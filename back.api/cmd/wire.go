package main

import (
	"main/internal/config"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
	"main/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func wireAuthDependency(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, cfg.Auth)
	authService := service.NewAuthService(userRepo, cfg.Auth)
	authHandler := handler.NewAuthHandler(authService, userUsecase)

	// Public 라우트
	public := r.Group("/api/v1/auth")
	{
		//1. 로그인
		public.POST("/login", authHandler.Login)
		//2. 회원가입
		public.POST("/register", authHandler.RegisterLocalUser)
		//3. 토큰 재발급
		public.POST("/refresh", authHandler.RefreshToken)
	}

	
}

func wirePostDependency(r *gin.Engine, authMiddleware gin.HandlerFunc, db *gorm.DB, cfg *config.Config) {
	postRepo := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUseCase(postRepo)
	postHandler := handler.NewPostHandler(postUsecase, *cfg)

	// Protected 라우트
	protected := r.Group("/api/v1/post").Use(authMiddleware)
	{
		protected.GET("/posts", postHandler.GetPosts)
		protected.POST("/posts", postHandler.CreatePost)
		protected.DELETE("/posts/:post_id", postHandler.DeletePost)
		protected.GET("/posts/:post_id", postHandler.GetPostDetail)
		protected.PATCH("/posts/:post_id", postHandler.UpdatePost)
	}
}

