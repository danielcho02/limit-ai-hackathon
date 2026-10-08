package handler

import (
	"main/internal/config"
	"main/internal/domain"
	"main/internal/errors"
	v1 "main/internal/handler/v1"
	"main/internal/presenter"
	"main/internal/usecase/interfaces"
	"main/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	uc interfaces.PostUsecase
	cfg config.Config
}

func NewPostHandler(uc interfaces.PostUsecase, cfg config.Config) *PostHandler {
	return &PostHandler{
		uc: uc,
		cfg: cfg,
	}
}


// CreatePost godoc
// @Summary      게시글 생성
// @Security     BearerAuth
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        request  body      v1.CreatePostRequest  true  "게시글 생성 요청 파라미터"
// @Success      201      {object}  v1.Response{data=v1.GetPostDetailResponse}
// @Failure      400      {object}  errors.AppError  "ErrShouldBindJson(5001): 잘못된 JSON 형식입니다. / ErrInvalidCategoryID(1003): 올바른 category_id 형태가 아닙니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      403      {object}  errors.AppError  "ErrForbidden(108): 작성자 본인 또는 Admin이 아닙니다. / ErrForbiddenNotice(108): 공지사항은 관리자만 작성할 수 있습니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999): 게시글 생성 실패"
// @Router       /api/v1/posts [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	// 게시글 생성 로직 구현
	ctx := c.Request.Context()


	user_id, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	
	var req v1.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	userRole, exists := c.Get("userRole")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}

	if (req.CategoryID == domain.NoticePost && userRole.(domain.RoleType) < domain.AdminRole) {
		c.Error(errors.ErrForbidden)
		return
	}

	post := &domain.Post{
		Title:      req.Title,
		Content:    req.Content,
		AuthorID:   user_id.(int),
		CategoryID: req.CategoryID,
	}
	createdPost, err := h.uc.CreatePost(ctx, post, req.Files)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, v1.Response{
		Message: "게시글 생성 성공",
		Data:    presenter.PresentPost(createdPost),
	})
}

// GetPosts godoc
// @Summary      게시글 목록 조회
// @Security     BearerAuth
// @Tags         Post
// @Produce      json
// @param        category_id query     int     false  "게시글 카테고리 ID (0: 공지, ??? 구체화 필요)"
// @param        title    query     string  false  "게시글 제목 검색 키워드"
// @param        author_id query     int     false  "작성자 ID"
// @param        offset	query     int     false  "페이지네이션 오프셋 (기본값: 0)"
// @param        limit	query     int     false  "페이지네이션 리밋 (기본값: 10)"
// @Success      200  {object}  v1.Response{data=v1.GetPostsResponse}
// @Failure      400  {object}  errors.AppError  "ErrInvalidCategoryID(1003): 올바른 category_id 형태가 아닙니다."
// @Failure      401  {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      500  {object}  errors.AppError  "ErrInternalServer(5000): 쿼리 파싱 실패 / ErrDatabase(9999): 목록 조회 실패"
// @Router       /api/v1/posts [get]
func (h *PostHandler) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()

	pq, err := utils.LoadQuery(c)
	if err != nil {
		c.Error(errors.ErrInternalServer.Wrap(err))
		return
	}

	postInfos, total, err := h.uc.GetPosts(ctx, pq)
	if err != nil {
		c.Error(err)
		return
	}

	pagination := utils.GetOffsetPagination(c.Request)
	c.JSON(200, v1.Response{
		Message: "게시글 목록 조회 성공",
		Data:	presenter.ToGetPostsResponse(postInfos, pagination, total)})
}

// GetPostDetail godoc
// @Summary      게시글 상세 조회
// @Security     BearerAuth
// @Tags         Post
// @Produce      json
// @Param        post_id  path     int  true  "조회할 게시글 ID"
// @Success      200      {object}  v1.Response{data=v1.GetPostDetailResponse}
// @Failure      400      {object}  errors.AppError  "ErrInvalidPostID(1002): post_id가 올바른 정수 형식이 아닙니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      404      {object}  errors.AppError  "ErrPostNotFound(1001): 존재하지 않는 게시글입니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999): 조회수 증가 또는 조회 실패"
// @Router       /api/v1/posts/{post_id} [get]
func (h *PostHandler) GetPostDetail(c *gin.Context) {
	ctx := c.Request.Context()

	strID := c.Param("post_id")
	post_id, err := strconv.Atoi(strID)
	if err != nil {
		c.Error(errors.ErrInvalidPostID.Wrap(err))
		return
	}

	post, err := h.uc.GetPostDetail(ctx, post_id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, v1.Response{
		Message: "게시글 상세 조회 성공",
		Data:    presenter.PresentPost(post),
	})
}

// UpdatePost godoc
// @Summary      게시글 수정
// @Description  작성자 본인 또는 Admin만 수정할 수 있습니다.
// @Security     BearerAuth
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        post_id  path      int                   true  "수정할 게시글 ID"
// @Param        request  body      v1.UpdatePostRequest  true  "게시글 수정 요청 파라미터"
// @Success      200      {object}  v1.Response{data=v1.GetPostDetailResponse}
// @Failure      400      {object}  errors.AppError  "ErrInvalidPostID(1002): post_id가 올바른 정수 형식이 아닙니다. / ErrShouldBindJson(5001): 잘못된 JSON 형식입니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      403      {object}  errors.AppError  "ErrForbidden(108): 작성자 본인 또는 Admin이 아닙니다."
// @Failure      404      {object}  errors.AppError  "ErrPostNotFound(1001): 존재하지 않는 게시글입니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999): 게시글 수정 실패"
// @Router       /api/v1/posts/{post_id} [patch]
func (h *PostHandler) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()

	postID := c.Param("post_id")
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.Error(errors.ErrInvalidPostID.Wrap(err))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	role, _ := c.Get("userRole")
	userRole, _ := role.(domain.RoleType)

	var req v1.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrShouldBindJson.Wrap(err))
		return
	}

	updatedPost, err := h.uc.UpdatePost(ctx, id, userID.(int), userRole, req.Title, req.Content, req.CategoryID, req.Files)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, v1.Response{
		Message: "게시글 수정 성공",
		Data:    presenter.PresentPost(updatedPost),
	})
}

// DeletePost godoc
// @Summary      게시글 삭제
// @Description  작성자 본인 또는 Admin만 삭제할 수 있습니다.
// @Security     BearerAuth
// @Tags         Post
// @Produce      json
// @Param        post_id  path     int  true  "삭제할 게시글 ID"
// @Success      204      {object}  v1.Response
// @Failure      400      {object}  errors.AppError  "ErrInvalidPostID(1002): post_id가 올바른 정수 형식이 아닙니다."
// @Failure      401      {object}  errors.AppError  "ErrTokenMissing(105) / ErrInvalidAuthHeader(104): Authorization 헤더가 없거나 형식이 올바르지 않습니다. / 유효하지 않은 토큰입니다."
// @Failure      403      {object}  errors.AppError  "ErrForbidden(108): 작성자 본인 또는 Admin이 아닙니다."
// @Failure      404      {object}  errors.AppError  "ErrPostNotFound(1001): 존재하지 않는 게시글입니다."
// @Failure      500      {object}  errors.AppError  "ErrDatabase(9999): 게시글 삭제 실패"
// @Router       /api/v1/posts/{post_id} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()

	postID := c.Param("post_id")
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.Error(errors.ErrInvalidPostID.Wrap(err))
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.Error(errors.ErrUnauthorized)
		return
	}
	role, _ := c.Get("userRole")
	userRole, _ := role.(domain.RoleType)

	if err := h.uc.DeletePost(ctx, id, userID.(int), userRole); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, v1.Response{
		Message: "게시글 삭제 성공",
	})
}