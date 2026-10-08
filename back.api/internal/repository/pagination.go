package repository

import (
	"main/internal/utils"

	"gorm.io/gorm"
)

func paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    if offset < 0 {
      offset = utils.DefaultOffset
    }
    if limit <= 0 {
      limit = utils.DefaultLimit
    } else if limit > utils.MaxLimit {
      limit = utils.MaxLimit
    }

    return db.Offset(offset).Limit(limit)
  }
}

// 다음과 같이 사용
// db.Scopes(Paginate(pq.Offset, pq.Limit)).Find(&users)