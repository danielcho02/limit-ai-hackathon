package utils

import (
	"errors"
	"main/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoadQuery(c *gin.Context) (domain.PostQuery, error) {
	var query domain.PostQuery

	title := c.Query("title")
	categoryID := c.Query("category_id")
	authorID := c.Query("author_id")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	if title != "" {
		query.Title = &title
	}

	if categoryID != "" {
		catID, err := strconv.Atoi(categoryID)
		if err != nil {
			return query, errors.New("invalid category_id")
		}
		query.CategoryID = &catID
	}

	if authorID != "" {
		authID, err := strconv.Atoi(authorID)
		if err != nil {
			return query, errors.New("invalid author_id")
		}
		query.AuthorID = &authID
	}

	if limit != "" {
		lmt, err := strconv.Atoi(limit)
		if err != nil {
			return query, errors.New("invalid limit")
		}
		query.Limit = lmt
	}

	if offset != "" {
		ofst, err := strconv.Atoi(offset)
		if err != nil {
			return query, errors.New("invalid offset")
		}
		query.Offset = ofst
	}

	return query, nil
}