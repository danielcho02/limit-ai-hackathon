package utils

import (
	"net/http"
	"strconv"
)

const (
	DefaultLimit  int = 10
	DefaultOffset int = 0

	MaxLimit      int = 100
)

type OffsetPagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type PaginationMetadata struct {
	TotalCount  int `json:"total_count"`
	CurrentPage int `json:"current_page"`
	NextOffset  *int `json:"next_offset,omitempty"`
	PrevOffset  *int `json:"prev_offset,omitempty"`
}

// NewOffsetPagination은 검증된 OffsetPagination을 생성하는 생성자
// Limit=0으로 인한 division by zero 방지 및 음수 offset 방지
func NewOffsetPagination(limit, offset int) OffsetPagination {
	if limit <= 0 {
		limit = DefaultLimit
	} else if limit > MaxLimit {
		limit = MaxLimit
	}

	if offset < 0 {
		offset = DefaultOffset
	}

	return OffsetPagination{Limit: limit, Offset: offset}
}

// GetOffsetPagination은 HTTP 요청의 쿼리 파라미터에서 페이지네이션 값을 파싱
func GetOffsetPagination(r *http.Request) OffsetPagination {
	q := r.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))

	return NewOffsetPagination(limit, offset)
}

// GetTotalPages는 전체 아이템 수를 기반으로 총 페이지 수를 반환
func (p OffsetPagination) GetTotalPages(totalItems int) int {
	if totalItems <= 0 {
		return 1
	}
	return (totalItems + p.Limit - 1) / p.Limit
}

// GetCurrentPage는 현재 페이지 번호를 반환
// offset이 totalItems를 초과하는 경우 마지막 페이지를 반환
func (p OffsetPagination) GetCurrentPage(totalItems int) int {
	totalPages := p.GetTotalPages(totalItems)
	currentPage := (p.Offset / p.Limit) + 1
	if currentPage > totalPages {
		return totalPages
	}
	return currentPage
}

// GetMetadata는 페이지네이션 응답에 필요한 메타데이터를 생성
func (p OffsetPagination) GetMetadata(totalItems int) PaginationMetadata {
	totalPages := p.GetTotalPages(totalItems)
	currentPage := p.GetCurrentPage(totalItems)

	hasNext := currentPage < totalPages
	hasPrev := currentPage > 1

	var nextOffset *int
	if hasNext {
		next := p.getNextOffset()
		nextOffset = &next
	}

	var prevOffset *int
	if hasPrev {
		prev := p.getPreviousOffset()
		prevOffset = &prev
	}

	return PaginationMetadata{
		TotalCount:  totalItems,
		CurrentPage: currentPage,
		NextOffset:  nextOffset,
		PrevOffset:  prevOffset,
	}
}

// GetLimitOffset은 limit과 offset 값을 반환
func (p OffsetPagination) getLimitOffset() (int, int) {
	return p.Limit, p.Offset
}

// GetNextOffset은 다음 페이지의 offset을 반환
func (p OffsetPagination) getNextOffset() int {
	return p.Offset + p.Limit
}

// GetPreviousOffset은 이전 페이지의 offset을 반환
func (p OffsetPagination) getPreviousOffset() int {
	prevOffset := p.Offset - p.Limit
	if prevOffset < 0 {
		return 0
	}
	return prevOffset
}