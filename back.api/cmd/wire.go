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

func wireUserDependency(r *gin.Engine, authMiddleware gin.HandlerFunc, db *gorm.DB) {
	userHandler := handler.NewUserHandler(db)

	protected := r.Group("/api/v1/user").Use(authMiddleware)
	{
		protected.PATCH("/profile", userHandler.UpdateProfile)
	}
}

func wirePostDependency(r *gin.Engine, authMiddleware gin.HandlerFunc, db *gorm.DB, cfg *config.Config) {
	postRepo := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUseCase(postRepo)
	postHandler := handler.NewPostHandler(postUsecase, *cfg)

	// Protected 라우트
	protected := r.Group("/api/v1/posts").Use(authMiddleware)
	{
		protected.GET("", postHandler.GetPosts)
		protected.POST("", postHandler.CreatePost)
		protected.DELETE("/:post_id", postHandler.DeletePost)
		protected.GET("/:post_id", postHandler.GetPostDetail)
		protected.PATCH("/:post_id", postHandler.UpdatePost)
	}
}

func wireFileDependency(r *gin.Engine, authMiddleware gin.HandlerFunc, db *gorm.DB, cfg *config.Config) {
	fileSvc := service.NewFileService(cfg.Storage)
	fileRepo := repository.NewFileRepository(db)
	fileUc := usecase.NewFileUsecase(fileSvc, fileRepo, cfg.Storage)
	fileHandler := handler.NewFileHandler(fileSvc, fileUc, fileRepo)

	protected := r.Group("/api/v1/files").Use(authMiddleware)
	{
		protected.POST("", fileHandler.UploadFile)
		protected.GET("/:file_id", fileHandler.DownloadFile)
		protected.DELETE("/:file_id", fileHandler.DeleteFile)
	}
}

func wireCommentDependency(r *gin.Engine, authMiddleware gin.HandlerFunc, db *gorm.DB, cfg *config.Config) {
	commentRepo := repository.NewCommentRepository(db)
	commentUc := usecase.NewCommentUsecase(commentRepo)
	commentHandler := handler.NewCommentHandler(commentUc)

	protected := r.Group("/api/v1/comments").Use(authMiddleware)
	{
		protected.POST("/:post_id", commentHandler.CreateComment)
		protected.DELETE("/:comment_id", commentHandler.DeleteComment)
	}
}